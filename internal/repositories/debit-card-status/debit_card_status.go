package debitcardstatus_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type DebitCardStatusRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.DebitCardStatus) error
}

type debitCardStatusRepo struct {
	db *sqlx.DB
}

func NewDebitCardStatusRepo(db *sqlx.DB) DebitCardStatusRepo {
	return &debitCardStatusRepo{
		db: db,
	}
}

func (repo *debitCardStatusRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.debit_card_status (
		card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		status VARCHAR(20) NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index debit_card_status_user_id_idx on public.debit_card_status using  btree (user_id);
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

func (repo *debitCardStatusRepo) Insert(tx *sql.Tx, req orm.DebitCardStatus) error {
	params := make([]interface{}, 7)
	params[0] = req.CardID
	params[1] = req.UserID
	params[2] = req.Status.NullString
	params[3] = req.CreatedBy
	params[4] = req.CreatedDate
	params[5] = req.UpdatedBy.NullString
	params[6] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.debit_card_status`
	col := `(
		card_id,
		user_id,
		status,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}
