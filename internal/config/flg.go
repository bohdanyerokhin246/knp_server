package config

import "time"

type Patient struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FullName    string `json:"fullName"`
	TherapistID uint   `json:"therapistID"`
}

type Therapist struct {
	ID       uint   `json:"ID,omitempty"`
	FullName string `json:"full_name,omitempty"`
}

type Diagnose struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Diagnose string `json:"diagnose"`
}

type Exam struct {
	ExamID      uint      `json:"examID" gorm:"primaryKey"`
	ExamDate    time.Time `json:"examDate"`
	PatientID   uint      `json:"patientID"`
	DiagnoseID  uint      `json:"diagnoseID"`
	TherapistID uint      `json:"therapistID"`
}

type ExamDetails struct {
	ExamID        uint      `json:"examID" gorm:"primaryKey"`
	ExamDate      time.Time `json:"exam_date"`
	FullName      string    `json:"full_name"`
	DiagnoseName  string    `json:"diagnose_name"`
	TherapistName string    `json:"therapist_name"`
}
