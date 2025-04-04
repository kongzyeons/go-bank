package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type Account struct {
	AccountID     string              `db:"account_id"`
	UserID        string              `db:"user_id"`
	Name          types.SQLNullString `db:"name"`
	Type          types.SQLNullString `db:"type"`
	Currency      types.SQLNullString `db:"currency"`
	AccountNumber types.SQLNullString `db:"account_number"`
	Issuer        types.SQLNullString `db:"issuer"`
	CreatedBy     string              `db:"created_by"`
	CreatedDate   time.Time           `db:"created_date"`
	UpdatedBy     types.SQLNullString `db:"updated_by"`
	UpdatedDate   types.SQLNullTime   `db:"updated_date"`
}
