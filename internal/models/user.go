package models

type UserSetGreetingReq struct {
	Greeting string `json:"greeting" example:"Have a nice day Clare" validate:"required,max=255"`
	UserID   string `json:"-" validate:"required"`
}
