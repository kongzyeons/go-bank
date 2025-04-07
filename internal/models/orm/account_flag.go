package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type AccountFlag struct {
	flagID      int64               `db:"flag_id"`
	AccountID   string              `db:"account_id"`
	UserID      string              `db:"user_id"`
	FlagType    string              `db:"flag_type"`
	FlagValue   string              `db:"flag_value"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
