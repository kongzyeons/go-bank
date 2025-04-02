package response

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`

	Data *T `json:"data"`

	ValidationErrors map[string]string `json:"validationErrors"`
}

func (r Response[T]) JSON(c *fiber.Ctx) error {
	return c.Status(r.StatusCode).JSON(r)
}

func Ok[T any](data *T) Response[T] {
	return Response[T]{
		Success:    true,
		StatusCode: 200,
		Data:       data,
	}
}

func BadRequest[T any](msgs ...string) Response[T] {
	message := http.StatusText(http.StatusBadRequest)
	if len(msgs) > 0 {
		message = strings.Join(msgs, " ")
	}
	log.Println(message)
	return Response[T]{
		Success:    false,
		StatusCode: 400,
		Message:    message,
	}
}

func ValidationFailed[T any](valMap map[string]string) Response[T] {
	message := "Validate Bad Request"
	jsonStr, err := json.Marshal(valMap)
	if err == nil {
		message = string(jsonStr)
	}
	log.Println(message)

	return Response[T]{
		Success:          false,
		StatusCode:       400,
		Message:          "Validate Bad Request",
		ValidationErrors: valMap,
	}
}

func Unauthorized[T any](msgs ...string) Response[T] {
	message := http.StatusText(http.StatusUnauthorized)
	if len(msgs) > 0 {
		message = strings.Join(msgs, " ")
	}
	log.Println(message)
	return Response[T]{
		Success:    false,
		StatusCode: 401,
		Message:    message,
	}
}

func Notfound[T any](msgs ...string) Response[T] {
	message := http.StatusText(http.StatusNotFound)
	if len(msgs) > 0 {
		message = strings.Join(msgs, " ")
	}
	log.Println(message)

	return Response[T]{
		Success:    false,
		StatusCode: 404,
		Message:    message,
	}
}

func InternalServerError[T any](err error, message string) Response[T] {
	log.Println(message)
	return Response[T]{
		Success:    false,
		StatusCode: 500,
		Message:    message,
	}
}

func ErrorWithCode[T any](err error, message string, statusCode int) Response[T] {
	log.Println(message)
	return Response[T]{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
	}
}
