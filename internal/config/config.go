package config

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Body       string `json:"body"`
	CreatedBy  string `json:"createdBy"`
	IsActual   bool   `json:"isActual"`
	CreateDate string `json:"createDate"`
	Link       string `json:"link"`
}
