package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	user_svc "github.com/kongzyeons/go-bank/internal/services/api/user"
)

type UserHandler interface {
	GetGeeting(c *fiber.Ctx) error
}

type userHandler struct {
	userSvc user_svc.UserSvc
}

func NewUserHandler(userSvc user_svc.UserSvc) UserHandler {
	return &userHandler{
		userSvc: userSvc,
	}
}

// GetGeeting godoc
// @summary GetGeeting
// @description GetGeeting
// @tags HomePage API
// @security ApiKeyAuth
// @id UserGetGeeting
// @accept json
// @produce json
// @Router /api/v1/user/greeting [get]
func (h *userHandler) GetGeeting(c *fiber.Ctx) error {
	req := models.UserGetGeetingReq{
		UserID:   fmt.Sprintf("%v", c.Locals("user_id")),
		Username: fmt.Sprintf("%v", c.Locals("username")),
	}
	res := h.userSvc.GetGeeting(req)
	return res.JSON(c)

}
