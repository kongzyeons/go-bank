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
	user_svc "github.com/kongzyeons/go-bank/internal/services/api/user"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestGetGeeting(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.UserGetGeetingRes]
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

			svc := user_svc.NewUserSvcMock()

			svc.On("GetGeeting", mock.Anything).Return(
				tc.res,
			)

			h := NewUserHandler(svc)

			app := fiber.New()
			app.Post("/", h.GetGeeting)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}
