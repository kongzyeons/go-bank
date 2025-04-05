package usergreeting_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type UserGreetingRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.UserGreeting) error

	GetByID(id string) (res *orm.UserGreeting, err error)
}

type uerGreetingRepo struct {
	db *sqlx.DB
}

func NewUserGreetingRepo(db *sqlx.DB) UserGreetingRepo {
	return &uerGreetingRepo{
		db: db,
	}
}

func (repo *uerGreetingRepo) CreateTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS public.user_greetings (
			user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			greeting VARCHAR(255) NOT NULL,
			created_by VARCHAR(100) NOT NULL,
			created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
			updated_by VARCHAR(100) NULL,
			updated_date TIMESTAMPTZ NULL
		);
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
func (repo *uerGreetingRepo) Insert(tx *sql.Tx, req orm.UserGreeting) error {

	params := make([]interface{}, 6)
	params[0] = req.UserID
	params[1] = req.Greeting
	params[2] = req.CreatedBy
	params[3] = req.CreatedDate
	params[4] = req.UpdatedBy.NullString
	params[5] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.user_greetings`
	col := `(
		user_id,
		greeting, 
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}

func (repo *uerGreetingRepo) GetByID(id string) (res *orm.UserGreeting, err error) {
	params := make([]interface{}, 1)
	params[0] = id

	sl := `SELECT *`
	from := `FROM public.user_greetings`
	condition := `WHERE user_id = $1`
	query := fmt.Sprintf("%s %s %s", sl, from, condition)

	var result orm.UserGreeting
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, nil
}
