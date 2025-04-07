package account_svc

import (
	"strings"
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/utils/types"
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
		CreatedBy:      req.CreatedBy.String,
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

func editToUpdate(req models.AccountEditReq, dataDB orm.AccountDetail) orm.AccountDetail {
	timeNow := time.Now().UTC()
	dataDB.Name = types.NewNullString(strings.TrimSpace(req.Name))
	dataDB.Color = types.NewNullString(strings.TrimSpace(req.Color))
	dataDB.UpdatedBy = types.NewNullString(strings.TrimSpace(req.Username))
	dataDB.UpdatedDate = types.NewNullTime(timeNow)
	return dataDB
}

func setIsmainToUpdate(req models.AccountSetIsmainReq, dataDB orm.AccountDetail, ismain bool) orm.AccountDetail {
	timeNow := time.Now().UTC()
	dataDB.IsManinAccount = ismain
	dataDB.UpdatedBy = types.NewNullString(strings.TrimSpace(req.Username))
	dataDB.UpdatedDate = types.NewNullTime(timeNow)
	return dataDB
}

func addMoneyToEvent(req models.AccountAddMoneyReq) events.AccountAddMoneyEvent {
	return events.AccountAddMoneyEvent{
		UserID:    strings.TrimSpace(req.UserID),
		Username:  strings.TrimSpace(req.Username),
		AccountID: strings.TrimSpace(req.AccountID),
		Ammount:   req.Ammount,
		Currency:  strings.TrimSpace(req.Currency),
	}
}

func withdrawlToEvent(req models.AccountWithdrawlReq) events.AccountWithldrawEvent {
	return events.AccountWithldrawEvent{
		UserID:    strings.TrimSpace(req.UserID),
		Username:  strings.TrimSpace(req.Username),
		AccountID: strings.TrimSpace(req.AccountID),
		Ammount:   req.Ammount,
		Currency:  strings.TrimSpace(req.Currency),
	}
}
