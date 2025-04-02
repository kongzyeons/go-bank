package models

type HomePageGetUserGreetingsReq struct {
	UserID string `json:"-"`
}

type HomePageGetUserGreetingsRes struct {
	Greeting string `json:"greeting"`
}
