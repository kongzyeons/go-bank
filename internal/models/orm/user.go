package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type User struct {
	UserID      string              `db:"user_id"`
	Name        string              `db:"name"`
	Password    types.SQLNullString `db:"password"`
	DummyCol1   types.SQLNullString `db:"dummy_col_1"`
	CreatedBy   types.SQLNullString `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
