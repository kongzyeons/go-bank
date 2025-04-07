package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type DebitCardDesign struct {
	CardID      string              `db:"card_id"`
	UserID      string              `db:"user_id"`
	Color       types.SQLNullString `db:"color"`
	BorderColor types.SQLNullString `db:"border_color"`
	DummyCol9   types.SQLNullString `db:"dummy_col_9"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
