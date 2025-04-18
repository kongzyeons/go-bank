package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type DebitCardVW struct {
	CardID      string              `db:"card_id"`
	UserID      string              `db:"user_id"`
	Name        types.SQLNullString `db:"name"`
	Status      types.SQLNullString `db:"status"`
	Number      types.SQLNullString `db:"number"`
	Issuer      types.SQLNullString `db:"issuer"`
	Color       types.SQLNullString `db:"color"`
	BorderColor types.SQLNullString `db:"border_color"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
