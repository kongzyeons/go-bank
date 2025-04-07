package debitcarddetails_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type DebitCardSDetailRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.DebitCardDetail) error
}

type debitCardSDetailRepo struct {
	db *sqlx.DB
}

func NewDebitCardSDetailRepo(db *sqlx.DB) DebitCardSDetailRepo {
	return &debitCardSDetailRepo{
		db: db,
	}
}

func (repo *debitCardSDetailRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.debit_card_details (
		card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		issuer VARCHAR(100) NULL,
		number VARCHAR(25) NULL,
		dummy_col_10 varchar(255) DEFAULT NULL,
		created_by VARCHAR(100) NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index debit_card_details_user_id_idx on public.debit_card_details using  btree (user_id);
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

func (repo *debitCardSDetailRepo) Insert(tx *sql.Tx, req orm.DebitCardDetail) error {
	params := make([]interface{}, 8)
	params[0] = req.CardID
	params[1] = req.UserID
	params[2] = req.Issuer.NullString
	params[3] = req.Number.NullString
	params[4] = req.CreatedBy.NullString
	params[5] = req.CreatedDate
	params[6] = req.UpdatedBy.NullString
	params[7] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.debit_card_details`
	col := `(
		card_id,
		user_id,
		issuer,
		number,
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
