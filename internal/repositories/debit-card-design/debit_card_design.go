package debitcarddesign_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type DebitCarddesignRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.DebitCardDesign) error
}

type debitCarddesignRepo struct {
	db *sqlx.DB
}

func NewDebitCarddesignRepo(db *sqlx.DB) DebitCarddesignRepo {
	return &debitCarddesignRepo{
		db: db,
	}
}

func (repo *debitCarddesignRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.debit_card_design (
		card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		color VARCHAR(10) NULL,
		border_color VARCHAR(10) NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index debit_card_design_user_id_idx on public.debit_card_design using  btree (user_id);
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

func (repo *debitCarddesignRepo) Insert(tx *sql.Tx, req orm.DebitCardDesign) error {
	params := make([]interface{}, 8)
	params[0] = req.CardID
	params[1] = req.UserID
	params[2] = req.Color.NullString
	params[3] = req.BorderColor.NullString
	params[4] = req.CreatedBy
	params[5] = req.CreatedDate
	params[6] = req.UpdatedBy.NullString
	params[7] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.debit_card_design`
	col := `(
		card_id,
		user_id,
		color,
		border_color,
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
