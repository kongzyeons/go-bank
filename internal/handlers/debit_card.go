package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	debitcard_svc "github.com/kongzyeons/go-bank/internal/services/api/debit-card"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type DebitCardHandler interface {
	GetList(c *fiber.Ctx) error
}

type debitCardHandler struct {
	debitCardSvc debitcard_svc.DebitCardSvc
}

func NewDebitCardHandler(debitCardSvc debitcard_svc.DebitCardSvc) DebitCardHandler {
	return &debitCardHandler{
		debitCardSvc: debitCardSvc,
	}
}

// GetList godoc
// @summary GetList
// @description GetList
// @tags HomePage API
// @security ApiKeyAuth
// @id debitCardGetList
// @accept json
// @produce json
// @param DebitCardGetListReq body models.DebitCardGetListReq true "request body"
// @Router /api/v1/debitCard/getlist [post]
func (h *debitCardHandler) GetList(c *fiber.Ctx) error {
	var req models.DebitCardGetListReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	req.UserID = fmt.Sprintf("%v", c.Locals("user_id"))
	res := h.debitCardSvc.GetList(req)
	return res.JSON(c)
}
