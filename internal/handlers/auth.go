package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Ping(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type authHandler struct {
	authSvc auth_service.AuthSvc
}

func NewAuthHandler(authSvc auth_service.AuthSvc) AuthHandler {
	return &authHandler{
		authSvc: authSvc,
	}
}

// Register godoc
// @summary Register
// @description Register
// @tags Auth API
// @security ApiKeyAuth
// @id AuthRegister
// @accept json
// @produce json
// @param AuthRegisterReq body models.AuthRegisterReq true "request body"
// @Router /api/v1/register [post]
func (h *authHandler) Register(c *fiber.Ctx) error {
	var req models.AuthRegisterReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	res := h.authSvc.Register(req)
	return res.JSON(c)
}

// Login godoc
// @summary Login
// @description Login
// @tags Auth API
// @security ApiKeyAuth
// @id AuthLogin
// @accept json
// @produce json
// @param AuthLoginReq body models.AuthLoginReq true "request body"
// @Router /api/v1/login [post]
func (h *authHandler) Login(c *fiber.Ctx) error {
	var req models.AuthLoginReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	res := h.authSvc.Login(req)
	return res.JSON(c)
}

// Ping godoc
// @summary Ping
// @description Ping
// @tags Auth API
// @security ApiKeyAuth
// @id AuthPing
// @accept json
// @produce json
// @Router /api/v1/auth/ping [get]
func (h *authHandler) Ping(c *fiber.Ctx) error {
	req := models.AuthPingReq{
		UserID:   fmt.Sprintf("%v", c.Locals("user_id")),
		Username: fmt.Sprintf("%v", c.Locals("username")),
	}
	res := h.authSvc.Ping(req)
	return res.JSON(c)
}

// Refresh godoc
// @summary Refresh
// @description Refresh
// @tags Auth API
// @security ApiKeyAuth
// @id AuthRefresh
// @accept json
// @produce json
// @param AuthRefreshReq body models.AuthRefreshReq true "request body"
// @Router /api/v1/auth/refresh [post]
func (h *authHandler) Refresh(c *fiber.Ctx) error {
	var req models.AuthRefreshReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	req.UserID = fmt.Sprintf("%v", c.Locals("user_id"))
	req.Username = fmt.Sprintf("%v", c.Locals("username"))
	res := h.authSvc.Refresh(req)
	return res.JSON(c)
}

// Logout godoc
// @summary Logout
// @description Logout
// @tags Auth API
// @security ApiKeyAuth
// @id AuthLogout
// @accept json
// @produce json
// @Router /api/v1/auth/logout [post]
func (h *authHandler) Logout(c *fiber.Ctx) error {
	user_id := fmt.Sprintf("%v", c.Locals("user_id"))

	var req models.AuthLogoutReq
	req.UserID = user_id

	res := h.authSvc.Logout(req)
	return res.JSON(c)
}
