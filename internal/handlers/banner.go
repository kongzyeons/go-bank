package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kongzyeons/go-bank/internal/models"
	banner_svc "github.com/kongzyeons/go-bank/internal/services/api/banner"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type BannerHandler interface {
	GetList(c *fiber.Ctx) error
	GetListTest(c *fiber.Ctx) error
}

type bannerHandler struct {
	bannerSvc banner_svc.BannerSvc
}

func NewBannerHandler(bannerSvc banner_svc.BannerSvc) BannerHandler {
	return &bannerHandler{
		bannerSvc: bannerSvc,
	}
}

// GetList godoc
// @summary GetList
// @description GetList
// @tags HomePage API
// @security ApiKeyAuth
// @id BannerGetList
// @accept json
// @produce json
// @param BannerGetListReq body models.BannerGetListReq true "request body"
// @Router /api/v1/banner/getlist [post]
func (h *bannerHandler) GetList(c *fiber.Ctx) error {
	var req models.BannerGetListReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	req.UserID = fmt.Sprintf("%v", c.Locals("user_id"))
	res := h.bannerSvc.GetList(req)
	return res.JSON(c)
}

// GetListTest godoc
// @summary GetListTest
// @description GetListTest
// @tags Test API
// @security ApiKeyAuth
// @id BannerGetListTest
// @accept json
// @produce json
// @param BannerGetListReq body models.BannerGetListReq true "request body"
// @Router /api/test/banner/getlist [post]
func (h *bannerHandler) GetListTest(c *fiber.Ctx) error {
	var req models.BannerGetListReq
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any]().JSON(c)
	}
	var defaultUUID uuid.UUID // zero value
	req.UserID = defaultUUID.String()
	res := h.bannerSvc.GetList(req)
	return res.JSON(c)
}
