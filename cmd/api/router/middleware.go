package router

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

// func GetProtected(c *fiber.Ctx) error {
// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "Missing or invalid token",
// 		})
// 	}
// 	fmt.Println(token)
// 	return c.JSON(fiber.Map{"message": "You are authorized!"})
// }

type MiddlewareAuth interface {
	AuthRequired(c *fiber.Ctx) error
}

type middlewareAuth struct {
	redisClient *redis.Client
}

func NewMiddlewareAuth(redisClient *redis.Client) MiddlewareAuth {
	return &middlewareAuth{
		redisClient: redisClient,
	}
}

func (m *middlewareAuth) AuthRequired(c *fiber.Ctx) error {
	headerAuth := c.Get("Authorization")
	tokenString := strings.TrimPrefix(headerAuth, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return response.Unauthorized[any](err.Error()).JSON(c)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// key := "authSvc::userID"
		key := fmt.Sprintf("authSvc::%v", claims["user_id"])

		authTokenStr, err := m.redisClient.Get(context.Background(), key).Result()
		if err == redis.Nil {
			return response.Unauthorized[any]("not found user_id").JSON(c)
		}
		if err != nil {
			return response.Unauthorized[any](err.Error()).JSON(c)
		}

		var authToken orm.Auth
		err = json.Unmarshal([]byte(authTokenStr), &authToken)
		if err != nil {
			return response.Unauthorized[any](err.Error()).JSON(c)
		}
		if authToken.AccToken != tokenString {
			return response.Unauthorized[any]("invalid token").JSON(c)
		}

		c.Locals("user_id", claims["user_id"])
		c.Locals("username", claims["username"])
	}

	return c.Next()
}
