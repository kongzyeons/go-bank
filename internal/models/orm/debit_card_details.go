package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type DebitCardDetail struct {
	CardID      string              `db:"card_id"`
	UserID      string              `db:"user_id"`
	Issuer      types.SQLNullString `db:"issuer"`
	Number      types.SQLNullString `db:"number"`
	DummyCol10  types.SQLNullString `db:"dummy_col_10"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
