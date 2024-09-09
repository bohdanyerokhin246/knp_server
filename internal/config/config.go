package config

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostBody   string `json:"postBody"`
	CreatedBy  string `json:"createdBy"`
	IsActual   bool   `json:"isActual"`
	CreateDate string `json:"createDate"`
}
