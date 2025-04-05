package transaction_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type TransactionRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.Transaction) error
}

type transactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) TransactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (repo *transactionRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.transactions (
		transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		name VARCHAR(100) NULL,
		image VARCHAR(255) NULL,
		isBank bool DEFAULT false NOT NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
	);
	create index transactions_user_id_idx on public.transactions using  btree (user_id);
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

func (repo *transactionRepo) Insert(tx *sql.Tx, req orm.Transaction) error {
	params := make([]interface{}, 6)
	params[0] = req.UserID
	params[1] = req.Name
	params[2] = req.Image.String
	params[3] = req.IsBank
	params[4] = req.CreatedBy
	params[5] = req.CreatedDate

	tableInsert := `INSERT INTO public.transactions`
	col := `(
		user_id,
		name,
		image,
		isBank,
		created_by,
		created_date
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
