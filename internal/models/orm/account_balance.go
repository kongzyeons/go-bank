package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type AccountBalance struct {
	AccountID   string               `db:"account_id"`
	UserID      string               `db:"user_id"`
	Amount      types.SQLNullFloat64 `db:"amount"`
	DummyCol4   types.SQLNullString  `db:"dummy_col_4"`
	CreatedBy   types.SQLNullString  `db:"created_by"`
	CreatedDate time.Time            `db:"created_date"`
	UpdatedBy   types.SQLNullString  `db:"updated_by"`
	UpdatedDate types.SQLNullTime    `db:"updated_date"`
}
