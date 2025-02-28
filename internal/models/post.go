package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Body       string `json:"body,omitempty"`
	CreatedBy  string `json:"createdBy,omitempty"`
	IsActual   bool   `json:"isActual,omitempty"`
	CreateDate string `json:"createDate,omitempty"`
	Link       string `json:"link,omitempty"`
}

type PostRequest struct {
	Post
	FileData string `json:"fileData,omitempty"`
	FileName string `json:"fileName,omitempty"`
}
