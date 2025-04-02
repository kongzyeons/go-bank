package auth_service

import (
	"strings"
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/utils/types"
)

func toUserRepoInsert(req models.AuthRegisterReq) orm.User {
	timeNow := time.Now().UTC()
	return orm.User{
		Name:        strings.TrimSpace(req.Username),
		Password:    strings.TrimSpace(req.Password),
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	}
}

func toUserGreeting(userID string) orm.UserGreeting {
	timeNow := time.Now().UTC()
	return orm.UserGreeting{
		UserID:      userID,
		Greeting:    "Have a nice day Clare",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	}
}
