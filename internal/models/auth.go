package models

type AuthRegisterReq struct {
	Username string `json:"username" example:"admin" validate:"required,max=100"`
	Password string `json:"password" example:"123456" validate:"required,min=6,max=6"`
}

type AuthLoginReq struct {
	Username string `json:"username" example:"admin" validate:"required,max=100"`
	Password string `json:"password" example:"123456" validate:"required,min=6,max=6"`
}
type AuthLoginRes struct {
	AccToken string `json:"accToken"`
	RefToken string `json:"refToken"`
}

type AuthPingReq struct {
	UserID   string `json:"-"`
	Username string `json:"-"`
}
type AuthPingRes struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
}

type AuthRefreshReq struct {
	RefToken string `json:"refToken"`
	UserID   string `json:"-"`
	Username string `json:"-"`
}
type AuthRefreshRes struct {
	AccToken string `json:"accToken"`
	RefToken string `json:"refToken"`
}

type AuthLogoutReq struct {
	UserID string `json:"-"`
}
