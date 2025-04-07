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
		Password:    types.NewNullString(strings.TrimSpace(req.Password)),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	}
}

func toUserGreeting(userID string) orm.UserGreeting {
	timeNow := time.Now().UTC()
	return orm.UserGreeting{
		UserID:      userID,
		Greeting:    types.NewNullString("Have a nice day Clare"),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	}
}
