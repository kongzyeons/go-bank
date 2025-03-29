package models

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterRes struct {
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRes struct {
	AccToken string `json:"acc_token"`
	RefToken string `json:"ref_token"`
}

type PingReq struct {
	RefToken string `json:"ref_token"`
}
type PingRes struct {
	AccToken string `json:"acc_token"`
	RefToken string `json:"ref_token"`
}
