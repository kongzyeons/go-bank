package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kongzyeons/go-bank/internal/meta"
)

type GenTokenReq struct {
	UserID       string        `json:"userID"`
	Username     string        `json:"username"`
	TimeDulation time.Duration `json:"timeDulation"`
}

func GenToken(req GenTokenReq) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  req.UserID,
		"username": req.Username,
		"exp":      time.Now().Add(req.TimeDulation).Unix(),
	},
	).SignedString([]byte(meta.KEY))

	return token, err

}
