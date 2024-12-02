package config

import "gorm.io/gorm"

type EMZ struct {
	gorm.Model
	Year                     int     `json:"year,omitempty"`
	Month                    int     `json:"month,omitempty"`
	DoctorName               string  `json:"doctorName,omitempty"`
	DoctorUnit               string  `json:"doctorUnit,omitempty"`
	EmzID                    string  `gorm:"uniqueIndex"`
	PatientsID               string  `json:"patientsID,omitempty"`
	Package                  string  `json:"package,omitempty"`
	ServicesIncluded         string  `json:"servicesIncluded,omitempty"`
	InterventionsIncluded    string  `json:"interventionsIncluded,omitempty"`
	AmountOfIncludedServices int     `json:"amountOfIncludedServices,omitempty"`
	CostEstimated            int     `json:"costEstimated,omitempty"`
	PaymentActual            float64 `json:"paymentActual,omitempty"`
	EmzIncludedToStatistic   string  `json:"emzIncludedToStatistic,omitempty"`
	EmzProblem               string  `json:"emzProblem,omitempty"`
}

type StatisticPatient struct {
	gorm.Model
	Year          int     `json:"year,omitempty"`
	Month         int     `json:"month,omitempty"`
	PatientID     string  `json:"patientID,omitempty"`
	Package       string  `json:"package,omitempty"`
	PaymentActual float64 `json:"paymentActual,omitempty"`
}

type SummarizeStatistic struct {
	Month           int     `json:"month,omitempty"`
	Year            int     `json:"year,omitempty"`
	DoctorName      string  `json:"doctorName,omitempty"`
	DoctorUnit      string  `json:"unit,omitempty"`
	Package         string  `json:"package,omitempty"`
	CountPatientsID int64   `json:"countPatientID,omitempty"`
	UniquePatients  int64   `json:"uniquePatients,omitempty"`
	SumIncluded     float64 `json:"sumIncluded,omitempty"`
	AvgIncluded     float64 `json:"avgIncluded,omitempty"`
	CostEstimated   float64 `json:"costEstimated,omitempty"`
	PaymentActual   float64 `json:"paymentActual,omitempty"`
}

type Dynamic struct {
	gorm.Model
	Year                     int     `json:"year,omitempty"`
	Month                    int     `json:"month,omitempty"`
	DoctorName               string  `json:"doctorName,omitempty"`
	DoctorUnit               string  `json:"doctorUnit,omitempty"`
	EmzID                    string  `json:"emzID,omitempty"`
	PatientsID               string  `json:"patientsID,omitempty"`
	Package                  string  `json:"package,omitempty"`
	ServicesIncluded         string  `json:"servicesIncluded,omitempty"`
	InterventionsIncluded    string  `json:"interventionsIncluded,omitempty"`
	AmountOfIncludedServices int     `json:"amountOfIncludedServices,omitempty"`
	CostEstimated            int     `json:"costEstimated,omitempty"`
	PaymentActual            float64 `json:"paymentActual,omitempty"`
	EmzIncludedToStatistic   string  `json:"emzIncludedToStatistic,omitempty"`
	EmzProblem               string  `json:"emzProblem,omitempty"`
}
