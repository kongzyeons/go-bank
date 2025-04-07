package debitcard_svc

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
)

func toDebitCardGetListResult(req orm.DebitCardVW) models.DebitCardGetListResult {
	return models.DebitCardGetListResult{
		CardID:      req.CardID,
		UserID:      req.UserID,
		Name:        req.Name.String,
		Status:      req.Status.String,
		Number:      req.Number.String,
		Issuer:      req.Issuer.String,
		Color:       req.Issuer.String,
		BorderColor: req.BorderColor.String,
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
