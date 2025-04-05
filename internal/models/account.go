package models

import (
	"reflect"
	"time"
)

type AccountGetListReq struct {
	SearchText     string `json:"searchText" example:"search by type or name"`
	IsManinAccount *bool  `json:"isManinAccount" example:"true"`
	UserID         string `json:"-" validate:"required"`
	Page           int64  `json:"page" example:"1" validate:"gte=1"`
	PerPage        int64  `json:"perPage" example:"10" validate:"gte=1"`
	SortBy         struct {
		Field     string       `json:"field" example:"updatedDate"`
		FieldType reflect.Kind `json:"-"`
		Mode      string       `json:"mode" example:"desc"`
	} `json:"sortBy"`
}

type AccountGetListResult struct {
	AccountID      string     `db:"account_id" json:"accountId"`
	UserID         string     `db:"user_id" json:"userId"`
	IsManinAccount bool       `db:"is_main_account" json:"isMainAccount"`
	Name           string     `db:"name" json:"name"`
	Type           string     `db:"type" json:"type"`
	AccountNumber  string     `db:"account_number" json:"accountNumber"`
	Issuer         string     `db:"issuer" json:"issuer"`
	Amount         float64    `db:"amount" json:"amount"`
	Currency       string     `db:"currency" json:"currency"`
	Color          string     `db:"color" json:"color"`
	Progress       int64      `db:"progress" json:"progress"`
	CreatedBy      string     `db:"created_by" json:"createdBy"`
	CreatedDate    *time.Time `db:"created_date" json:"createdDate"`
	UpdatedBy      string     `db:"updated_by" json:"updatedBy"`
	UpdatedDate    *time.Time `db:"updated_date" json:"updatedDate"`
}

type AccountGetListRes struct {
	TotalPages   int64                  `json:"totalPages"`
	TotalResults int64                  `json:"totalResults"`
	Page         int64                  `json:"page"`
	PerPage      int64                  `json:"perPage"`
	TotalAnount  float64                `json:"totalAnount"`
	Results      []AccountGetListResult `json:"results"`
}

type AccountEditReq struct {
	AccountID string `json:"-" validate:"required"`
	UserID    string `json:"-" validate:"required"`
	Username  string `json:"-" validate:"required"`
	Name      string `json:"name" example:"name" validate:"required,max=100"`
	Color     string `json:"color" example:"color" validate:"required,max=10"`
}

type AccountEditRes struct {
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedDate time.Time `json:"updatedDate"`
}
