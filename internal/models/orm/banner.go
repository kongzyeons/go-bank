package orm

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type Banner struct {
	BannerID    string              `db:"banner_id"`
	UserID      string              `db:"user_id"`
	Title       types.SQLNullString `db:"title"`
	Description types.SQLNullString `db:"description"`
	Image       types.SQLNullString `db:"image"`
	CreatedBy   string              `db:"created_by"`
	CreatedDate time.Time           `db:"created_date"`
	UpdatedBy   types.SQLNullString `db:"updated_by"`
	UpdatedDate types.SQLNullTime   `db:"updated_date"`
}
