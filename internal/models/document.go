package models

import (
	"gorm.io/gorm"
	"time"
)

type Document struct {
	gorm.Model
	DocumentID  string    `json:"documentID,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty"`
	CreateDate  time.Time `json:"createDate,omitempty"`
	ExpiredDate time.Time `json:"expiredDate,omitempty"`
	WordLink    string    `json:"wordLink,omitempty"`
	ScanLink    string    `json:"scanLink,omitempty"`
}

type DocumentRequest struct {
	Document
	WordData string `json:"wordData,omitempty"`
	WordName string `json:"wordName,omitempty"`
	ScanData string `json:"scanData,omitempty"`
	ScanName string `json:"scanName,omitempty"`
}
