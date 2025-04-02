package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	homepage_svc "github.com/kongzyeons/go-bank/internal/services/api/homepage"
)

type HomePageHandler interface {
	GetUserGreetings(c *fiber.Ctx) error
}

type homePageHandler struct {
	homePageSvc homepage_svc.HomePageSvc
}

func NewHomePageHandler(homePageSvc homepage_svc.HomePageSvc) HomePageHandler {
	return &homePageHandler{
		homePageSvc: homePageSvc,
	}
}

// GetUserGreetings godoc
// @summary GetUserGreetings
// @description GetUserGreetings
// @tags HomePage API
// @security ApiKeyAuth
// @id HomePageGetUserGreetings
// @accept json
// @produce json
// @Router /api/v1/homepage/greeting [get]
func (h *homePageHandler) GetUserGreetings(c *fiber.Ctx) error {
	req := models.HomePageGetUserGreetingsReq{
		UserID: fmt.Sprintf("%v", c.Locals("user_id")),
	}
	res := h.homePageSvc.GetUserGreetings(req)
	return res.JSON(c)
}
