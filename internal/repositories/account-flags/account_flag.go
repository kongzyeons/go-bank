package accountflag_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountFlagRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.AccountFlag) error
}

type accountFlagRepo struct {
	db *sqlx.DB
}

func NewAccountFlagRepo(db *sqlx.DB) AccountFlagRepo {
	return &accountFlagRepo{
		db: db,
	}
}

func (repo *accountFlagRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.account_flags (
 		flag_id SERIAL PRIMARY KEY, 
		account_id UUID DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		flag_type VARCHAR(50) NOT NULL,
		flag_value VARCHAR(30) NOT NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index account_flags_account_id_idx on public.account_flags using  btree (account_id);
	create index account_flags_user_id_idx on public.account_flags using  btree (user_id);
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

func (repo *accountFlagRepo) Insert(tx *sql.Tx, req orm.AccountFlag) error {
	params := make([]interface{}, 8)
	params[0] = req.AccountID
	params[1] = req.UserID
	params[2] = req.FlagType
	params[3] = req.FlagValue
	params[4] = req.CreatedBy
	params[5] = req.CreatedDate
	params[6] = req.UpdatedBy.NullString
	params[7] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.account_flags`
	col := `(
		account_id,
		user_id,
		flag_type,
		flag_value,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}
