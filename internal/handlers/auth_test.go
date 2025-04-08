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
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[any]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AuthRegisterReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := auth_service.NewAuthSvcMock()

			svc.On("Register", mock.Anything).Return(
				tc.res,
			)

			h := NewAuthHandler(svc)

			app := fiber.New()
			app.Post("/", h.Register)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AuthLoginRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AuthRegisterReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := auth_service.NewAuthSvcMock()

			svc.On("Login", mock.Anything).Return(
				tc.res,
			)

			h := NewAuthHandler(svc)

			app := fiber.New()
			app.Post("/", h.Login)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestPing(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AuthPingRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AuthRegisterReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := auth_service.NewAuthSvcMock()

			svc.On("Ping", mock.Anything).Return(
				tc.res,
			)

			h := NewAuthHandler(svc)

			app := fiber.New()
			app.Post("/", h.Ping)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestRefresh(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.AuthRefreshRes]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AuthRegisterReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := auth_service.NewAuthSvcMock()

			svc.On("Refresh", mock.Anything).Return(
				tc.res,
			)

			h := NewAuthHandler(svc)

			app := fiber.New()
			app.Post("/", h.Refresh)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}

func TestLogout(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[any]
	}{
		{nameTest: "test"},
		{nameTest: "test",
			reqBodyBytes: func() io.Reader {
				jsonStr, _ := json.Marshal(models.AuthRegisterReq{})
				return bytes.NewReader(jsonStr)
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			svc := auth_service.NewAuthSvcMock()

			svc.On("Logout", mock.Anything).Return(
				tc.res,
			)

			h := NewAuthHandler(svc)

			app := fiber.New()
			app.Post("/", h.Logout)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}
