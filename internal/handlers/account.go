package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	account_svc "github.com/kongzyeons/go-bank/internal/services/api/account"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type AccountHandler interface {
	GetList(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
}

type accountHandler struct {
	accountSvc account_svc.AccountSvc
}

func NewAccountHandler(accountSvc account_svc.AccountSvc) AccountHandler {
	return &accountHandler{
		accountSvc: accountSvc,
	}
}

// GetList godoc
// @summary GetList
// @description GetList
// @tags HomePage API
// @security ApiKeyAuth
// @id AccountGetList
// @accept json
// @produce json
// @param AccountGetListReq body models.AccountGetListReq true "request body"
// @Router /api/v1/account/getlist [post]
func (h *accountHandler) GetList(c *fiber.Ctx) error {
	var req models.AccountGetListReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	req.UserID = fmt.Sprintf("%v", c.Locals("user_id"))
	res := h.accountSvc.GetList(req)
	return res.JSON(c)
}

// Edit godoc
// @summary Edit
// @description Edit
// @tags Account API
// @security ApiKeyAuth
// @id AccountEdit
// @accept json
// @produce json
// @param accountID path string true "accountID"
// @param AccountEditReq body models.AccountEditReq true "request body"
// @Router /api/v1/account/edit/{accountID} [put]
func (h *accountHandler) Edit(c *fiber.Ctx) error {
	var req models.AccountEditReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	req.AccountID = c.Params("accountID")
	req.UserID = fmt.Sprintf("%v", c.Locals("user_id"))
	req.Username = fmt.Sprintf("%v", c.Locals("username"))
	res := h.accountSvc.Edit(req)
	return res.JSON(c)
}
