package user_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type UserRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.User) (userID string, err error)
	InsertMock(tx *sql.Tx, req orm.User) (userID string, err error)

	GetByID(id string) (res *orm.User, err error)
	GetUnique(name string) (res *orm.User, err error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) CreateTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS public.users (
			user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(100) NULL,
			password VARCHAR(6) NULL,
			dummy_col_1 varchar(255) DEFAULT NULL,
			created_by VARCHAR(100) NULL,
			created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
			updated_by VARCHAR(100) NULL,
			updated_date TIMESTAMPTZ NULL
		);
		create index users_name_idx on public.users using  btree (name);
`
	tx, err := repo.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *userRepo) Insert(tx *sql.Tx, req orm.User) (userID string, err error) {
	params := make([]interface{}, 6)
	params[0] = req.Name
	params[1] = req.Password.NullString
	params[2] = req.CreatedBy.NullString
	params[3] = req.CreatedDate
	params[4] = req.UpdatedBy.NullString
	params[5] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.users`
	col := `(
		name,
		password, 
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6
	)`

	returning := `RETURNING user_id`

	query := fmt.Sprintf(`%s %s %s %s;`, tableInsert, col, values, returning)

	err = tx.QueryRow(query, params...).Scan(&userID)
	if postgresql.IsSQLReallyError(err) {
		return userID, err
	}

	return userID, nil
}

func (repo *userRepo) InsertMock(tx *sql.Tx, req orm.User) (userID string, err error) {
	var defaultUUID uuid.UUID // zero value
	params := make([]interface{}, 7)
	params[0] = defaultUUID.String()
	params[1] = req.Name
	params[2] = req.Password.NullString
	params[3] = req.CreatedBy.NullString
	params[4] = req.CreatedDate
	params[5] = req.UpdatedBy.NullString
	params[6] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.users`
	col := `(
		user_id,
		name,
		password, 
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7
	)`

	returning := `RETURNING user_id`

	query := fmt.Sprintf(`%s %s %s %s;`, tableInsert, col, values, returning)

	err = tx.QueryRow(query, params...).Scan(&userID)
	if postgresql.IsSQLReallyError(err) {
		return userID, err
	}

	return userID, nil
}

func (repo *userRepo) GetByID(id string) (res *orm.User, err error) {
	params := make([]interface{}, 1)
	params[0] = id

	sl := `SELECT *`
	from := `FROM public.users`
	condition := `WHERE user_id = $1`
	query := fmt.Sprintf("%s %s %s", sl, from, condition)

	var result orm.User
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, nil
}
func (repo *userRepo) GetUnique(name string) (res *orm.User, err error) {
	params := make([]interface{}, 1)
	params[0] = name

	sl := `SELECT *`
	from := `FROM public.users`
	condition := `WHERE name = $1`
	query := fmt.Sprintf("%s %s %s", sl, from, condition)

	var result orm.User
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, nil
}
