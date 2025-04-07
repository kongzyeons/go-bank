package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type Transaction struct {
	transactionID string              `db:"transaction_id"`
	UserID        string              `db:"user_id"`
	Name          types.SQLNullString `db:"name"`
	Image         types.SQLNullString `db:"image"`
	IsBank        bool                `db:"isBank"`
	DummyCol6     types.SQLNullString `db:"dummy_col_6"`
	CreatedBy     types.SQLNullString `db:"created_by"`
	CreatedDate   time.Time           `db:"created_date"`
}
