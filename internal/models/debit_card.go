package models

import (
	"reflect"
	"time"
)

type DebitCardGetListReq struct {
	SearchText string `json:"searchText" example:"search by name"`
	Status     string `json:"status" example:""`
	UserID     string `json:"-" validate:"required"`
	Page       int64  `json:"page" example:"1" validate:"gte=1"`
	PerPage    int64  `json:"perPage" example:"10" validate:"gte=1"`
	SortBy     struct {
		Field     string       `json:"field" example:"updatedDate"`
		FieldType reflect.Kind `json:"-"`
		Mode      string       `json:"mode" example:"desc"`
	} `json:"sortBy"`
}

type DebitCardGetListResult struct {
	CardID      string     `db:"card_id" json:"cardId"`
	UserID      string     `db:"user_id" json:"userId"`
	Name        string     `db:"name" json:"name"`
	Status      string     `db:"status" json:"status"`
	Number      string     `db:"number" json:"number"`
	Issuer      string     `db:"issuer" json:"issuer"`
	Color       string     `db:"color" json:"color"`
	BorderColor string     `db:"border_color" json:"borderColor"`
	CreatedBy   string     `db:"created_by" json:"createdBy"`
	CreatedDate *time.Time `db:"created_date" json:"createdDate"`
	UpdatedBy   string     `db:"updated_by" json:"updatedBy"`
	UpdatedDate *time.Time `db:"updated_date" json:"updatedDate"`
}

type DebitCardGetListRes struct {
	TotalPages   int64                    `json:"totalPages"`
	TotalResults int64                    `json:"totalResults"`
	Page         int64                    `json:"page"`
	PerPage      int64                    `json:"perPage"`
	Results      []DebitCardGetListResult `json:"results"`
}
