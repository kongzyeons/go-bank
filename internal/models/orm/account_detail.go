package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type AccountDetail struct {
	AccountID      string              `db:"account_id"`
	UserID         string              `db:"user_id"`
	Name           types.SQLNullString `db:"name"`
	Color          types.SQLNullString `db:"color"`
	IsManinAccount bool                `db:"is_main_account"`
	Progress       types.SQLNullInt64  `db:"progress"`
	DummyCol5      types.SQLNullString `db:"dummy_col_5"`
	CreatedBy      types.SQLNullString `db:"created_by"`
	CreatedDate    time.Time           `db:"created_date"`
	UpdatedBy      types.SQLNullString `db:"updated_by"`
	UpdatedDate    types.SQLNullTime   `db:"updated_date"`
}
