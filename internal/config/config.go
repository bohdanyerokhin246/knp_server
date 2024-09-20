package config

import "gorm.io/gorm"

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

type Statistic struct {
	gorm.Model
	FullName       string `json:"fullName,omitempty"`
	Unit           string `json:"unit,omitempty"`
	Package        string `json:"package,omitempty"`
	EmzAmount      int    `json:"emzAmount,omitempty"`
	PatientsAmount int    `json:"patientsAmount,omitempty"`
	ServicesAmount int    `json:"servicesAmount,omitempty"`
	Rate           int    `json:"rate,omitempty"`
	Payment        int    `json:"payment,omitempty"`
}

type StatisticFromExcel struct {
	gorm.Model
	DoctorName               string `json:"doctorName,omitempty"`
	EmzID                    string `json:"emzID,omitempty"`
	PatientsID               string `json:"patientsID,omitempty"`
	ServicesIncluded         string `json:"servicesIncluded,omitempty"`
	AmountOfIncludedServices int    `json:"amountOfIncludedServices,omitempty"`
	CostEstimated            int    `json:"costEstimated,omitempty"`
	PaymentActual            int    `json:"paymentActual,omitempty"`
	EmzIncludedToStatistic   string `json:"emzIncludedToStatistic,omitempty"`
	EmzProblem               string `json:"emzProblem,omitempty"`
}

type Patient struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Patronymic string `json:"patronymic"`
}

type Diagnose struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Diagnose string `json:"diagnose"`
}

type Exam struct {
	ExamID     uint `json:"examID" gorm:"primaryKey"`
	PatientID  uint `json:"patientID"`
	DiagnoseID uint `json:"diagnoseID"`
}
