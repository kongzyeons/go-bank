package banner_svc

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
)

func toBannerGetListResult(req orm.Banner) models.BannerGetListResult {
	return models.BannerGetListResult{
		BannerID:    req.BannerID,
		UserID:      req.UserID,
		Title:       req.Title.String,
		Description: req.Description.String,
		Image:       req.Image.String,
		CreatedBy:   req.CreatedBy.String,
		CreatedDate: func() *time.Time {
			if req.CreatedDate.IsZero() {
				return nil
			}
			createdDate := req.CreatedDate.UTC()
			return &createdDate
		}(),
		UpdatedBy: req.UpdatedBy.String,
		UpdatedDate: func() *time.Time {
			if req.UpdatedDate.IsNull() {
				return nil
			}
			updatedDate := req.UpdatedDate.Time.UTC()
			return &updatedDate
		}(),
	}
}
