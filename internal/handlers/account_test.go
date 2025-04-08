package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kongzyeons/go-bank/internal/models"
	account_svc "github.com/kongzyeons/go-bank/internal/services/api/account"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestGetList(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AccountGetListRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountGetListReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("GetList", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.GetList)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestEdit(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AccountEditRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountEditReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("Edit", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.Edit)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestGetQrcode(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AccountGetQrcodeRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountGetQrcodeReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("GetQrcode", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.GetQrcode)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestSetIsmain(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AccountSetIsmainRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountSetIsmainReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("SetIsmain", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.SetIsmain)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestAddMoney(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[any]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountAddMoneyReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("AddMoney", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.AddMoney)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestWithdrawl(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[any]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AccountAddMoneyReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := account_svc.NewAccountSvcMock()

			svc.On("Withdrawl", mock.Anything).Return(
				tc.res,
			)

			h := NewAccountHandler(svc)

			app := fiber.New()
			app.Post("/", h.Withdrawl)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}
