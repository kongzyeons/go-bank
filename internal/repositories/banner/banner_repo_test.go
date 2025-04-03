package banner_repo

import (
	"reflect"
	"testing"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestGetList(t *testing.T) {
	// init db
	db := postgresql.InitPostgresql()
	defer db.Close()

	repo := NewBannerRepo(db)
	repo.GetList(models.BannerGetListReq{
		SearchText: "search by title",
		Page:       1,
		PerPage:    10,
		SortBy: struct {
			Field     string       "json:\"field\" example:\"updatedDate\""
			FieldType reflect.Kind "json:\"-\""
			Mode      string       "json:\"mode\" example:\"desc\""
		}{
			Field: "updatedDate",
			Mode:  "asc",
		},
	})
}
