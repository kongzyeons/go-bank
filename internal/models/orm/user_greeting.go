package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type UserGreeting struct {
	UserID      string              `db:"user_id"`
	Greeting    types.SQLNullString `db:"greeting"`
	DummyCol2   types.SQLNullString `db:"dummy_col_2"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
