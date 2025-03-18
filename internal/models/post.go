package models

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Body       string   `json:"body,omitempty"`
	CreatedBy  string   `json:"createdBy,omitempty"`
	IsActual   bool     `json:"isActual,omitempty"`
	Link       string   `json:"link,omitempty"`
	FileInfoID uint     `json:"fileInfoID,omitempty"`
	FileInfo   FileInfo `json:"fileInfo,omitempty"`
}

type NewsRequest struct {
}

type FileInfo struct {
	gorm.Model
	FileData string `json:"fileData,omitempty"`
	FileName string `json:"fileName,omitempty"`
}
