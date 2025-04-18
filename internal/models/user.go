package models

type UserGetGeetingReq struct {
	UserID   string `json:"-" validate:"required"`
	Username string `json:"-" validate:"required"`
}

type UserGetGeetingRes struct {
	Username string `json:"name"`
	Greeting string `json:"greeting"`
}

type UserSetGreetingReq struct {
	Greeting string `json:"greeting" example:"Have a nice day Clare" validate:"required,max=255"`
	UserID   string `json:"-" validate:"required"`
}
