package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type DebitCardStatus struct {
	CardID      string              `db:"card_id"`
	UserID      string              `db:"user_id"`
	Status      types.SQLNullString `db:"status"`
	DummyCol8   types.SQLNullString `db:"dummy_col_8"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
