package account_svc

import (
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
)

func toBannerGetListResult(req orm.AccountVW) models.AccountGetListResult {
	return models.AccountGetListResult{
		AccountID:      req.AccountID,
		UserID:         req.UserID,
		IsManinAccount: req.IsManinAccount,
		Name:           req.Name.String,
		Type:           req.Type.String,
		AccountNumber:  req.AccountNumber.String,
		Issuer:         req.Issuer.String,
		Amount:         float64(int(req.Amount.Float64*100)) / 100,
		Currency:       req.Currency.String,
		Color:          req.Color.String,
		Progress:       req.Progress.Int64,
		CreatedBy:      req.CreatedBy,
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
