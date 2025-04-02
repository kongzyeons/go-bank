package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type UserGreeting struct {
	UserID      string              `db:"user_id"`
	Greeting    string              `db:"greeting"`
	CreatedBy   string              `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
