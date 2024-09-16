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

type PostRequest struct {
	Post
	FileData string `json:"fileData"`
	FileName string `json:"fileName"`
}

type Statistic struct {
	gorm.Model
	FullName       string `json:"fullName"`
	Unit           string `json:"unit"`
	Package        string `json:"package"`
	EmzAmount      int    `json:"emzAmount"`
	PatientsAmount int    `json:"patientsAmount"`
	ServicesAmount int    `json:"servicesAmount"`
	Rate           int    `json:"rate"`
	Payment        int    `json:"payment"`
}
