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
	debitcard_svc "github.com/kongzyeons/go-bank/internal/services/api/debit-card"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestGetListCard(t *testing.T) {
	testCases := []struct {
		nameTest     string
		reqBodyBytes io.Reader
		res          response.Response[*models.DebitCardGetListRes]
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

			svc := debitcard_svc.NewDebitCardSvcMock()

			svc.On("GetList", mock.Anything).Return(
				tc.res,
			)

			h := NewDebitCardHandler(svc)

			app := fiber.New()
			app.Post("/", h.GetList)

			req := httptest.NewRequest(http.MethodPost, "/", tc.reqBodyBytes)
			req.Header.Set("Content-Type", "application/json")
			app.Test(req)

		})
	}
}
