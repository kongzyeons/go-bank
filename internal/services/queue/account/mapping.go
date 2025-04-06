package account_queue

import (
	"strings"
	"time"

	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/utils/types"
)

func addMoneyToUpdate(req events.AccountAddMoneyEvent, dataDB orm.AccountBalance) orm.AccountBalance {
	timeNow := time.Now().UTC()
	dataDB.Amount = types.NewNullFloat64(dataDB.Amount.Float64 + req.Ammount)
	dataDB.UpdatedBy = types.NewNullString(strings.TrimSpace(req.Username))
	dataDB.UpdatedDate = types.NewNullTime(timeNow)
	return dataDB
}

func withdrawlToUpdate(req events.AccountWithldrawEvent, dataDB orm.AccountBalance) orm.AccountBalance {
	timeNow := time.Now().UTC()
	dataDB.Amount = types.NewNullFloat64(dataDB.Amount.Float64 - req.Ammount)
	dataDB.UpdatedBy = types.NewNullString(strings.TrimSpace(req.Username))
	dataDB.UpdatedDate = types.NewNullTime(timeNow)
	return dataDB
}
