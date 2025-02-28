package models

import (
	"gorm.io/gorm"
	"time"
)

type Kep struct {
	gorm.Model
	FullName         string    `json:"fullName,omitempty"`
	RNOKPP           string    `json:"RNOKPP,omitempty"`
	ZNOKId           int       `json:"ZNOKId,omitempty"`
	ECPId            int       `json:"ECPId,omitempty"`
	ECPCreateDate    time.Time `json:"ECPCreateDate"`
	ECPExpiredDate   time.Time `json:"ECPExpiredDate"`
	ProtoId          int       `json:"protoId,omitempty"`
	ProtoCreateDate  time.Time `json:"protoCreateDate"`
	ProtoExpiredDate time.Time `json:"protoExpiredDate"`
}
