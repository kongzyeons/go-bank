package accountdetail_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountDetailRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.AccountDetail) error
}

type accountDetailRepo struct {
	db *sqlx.DB
}

func NewAccountDetailRepo(db *sqlx.DB) AccountDetailRepo {
	return &accountDetailRepo{
		db: db,
	}
}

func (repo *accountDetailRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.account_details (
		account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		color VARCHAR(10) NULL,
		is_main_account bool DEFAULT false NOT NULL,
		progress BIGINT NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index account_details_user_id_idx on public.account_details using  btree (user_id);
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

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (repo *accountDetailRepo) Insert(tx *sql.Tx, req orm.AccountDetail) error {
	params := make([]interface{}, 9)
	params[0] = req.AccountID
	params[1] = req.UserID
	params[2] = req.Color.NullString
	params[3] = req.IsManinAccount
	params[4] = req.Progress.NullInt64
	params[5] = req.CreatedBy
	params[6] = req.CreatedDate
	params[7] = req.UpdatedBy.NullString
	params[8] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.account_details`
	col := `(
		account_id,
		user_id,
		color,
		is_main_account,
		progress,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}
