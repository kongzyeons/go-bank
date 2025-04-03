package models

import (
	"reflect"
	"time"
)

type BannerGetListReq struct {
	SearchText string `json:"searchText" example:"search by title"`
	UserID     string `json:"-" validate:"required"`
	Page       int64  `json:"page" example:"1" validate:"gte=1"`
	PerPage    int64  `json:"perPage" example:"10" validate:"gte=1"`
	SortBy     struct {
		Field     string       `json:"field" example:"updatedDate"`
		FieldType reflect.Kind `json:"-"`
		Mode      string       `json:"mode" example:"desc"`
	} `json:"sortBy"`
}

type BannerGetListResult struct {
	BannerID    string     `json:"bannerID" db:"banner_id"`
	UserID      string     `json:"userID" db:"user_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Image       string     `json:"image" db:"image"`
	CreatedBy   string     `json:"createdBy" db:"created_by"`
	CreatedDate *time.Time `json:"createdDate" db:"created_date"`
	UpdatedBy   string     `json:"updatedBy" db:"updated_by"`
	UpdatedDate *time.Time `json:"updatedDate" db:"updated_date"`
}

type BannerGetListRes struct {
	TotalPages   int64                 `json:"totalPages"`
	TotalResults int64                 `json:"totalResults"`
	Page         int64                 `json:"page"`
	PerPage      int64                 `json:"perPage"`
	Results      []BannerGetListResult `json:"results"`
}
