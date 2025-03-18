package models

type ODK struct {
	ID        uint       `gorm:"primaryKey" json:"id,omitempty"`
	Name      string     `gorm:"unique;not null" json:"name,omitempty"`
	Diagnoses []Diagnose `gorm:"foreignKey:ODKID;constraint:OnDelete:CASCADE;" json:"diagnoses,omitempty"`
}

type Diagnose struct {
	ID    uint   `gorm:"primaryKey" json:"id,omitempty"`
	ODKID string `gorm:"not null" json:"odk_id,omitempty"` // Внешний ключ
	Code  string `gorm:"not null" json:"code,omitempty"`
	Name  string `gorm:"not null" json:"name,omitempty"`
}

type Specialist struct {
	ID   uint   `gorm:"primaryKey" json:"id,omitempty"`
	Code string `gorm:"unique;not null" json:"code,omitempty"`
	Name string `gorm:"not null" json:"name,omitempty"`
}

type Procedure struct {
	ID               uint         `gorm:"primaryKey" json:"id,omitempty"`
	ClassNumber      int          `gorm:"not null" json:"class_number,omitempty"`
	Class            string       `gorm:"not null" json:"class,omitempty"`
	Code             string       `gorm:"unique;not null" json:"code,omitempty"`
	InterventionName string       `gorm:"not null" json:"intervention_name,omitempty"`
	ODKs             []ODK        `gorm:"many2many:procedure_odks;" json:"odks,omitempty"`
	Specialists      []Specialist `gorm:"many2many:procedure_specialists;" json:"specialists,omitempty"`
}

type ProcedureODK struct {
	ProcedureID uint `gorm:"primaryKey" json:"procedure_id,omitempty"`
	ODKID       uint `gorm:"primaryKey" json:"odk_id,omitempty"`
}

type ProcedureSpecialist struct {
	ProcedureID  uint `gorm:"primaryKey" json:"procedure_id,omitempty"`
	SpecialistID uint `gorm:"primaryKey" json:"specialist_id,omitempty"`
}

type Consultation struct {
	ID               uint         `gorm:"primaryKey" json:"id,omitempty"`
	ClassNumber      int          `gorm:"not null" json:"class_number,omitempty"`
	ClassName        string       `gorm:"not null" json:"className,omitempty"`
	Code             string       `gorm:"not null" json:"code,omitempty"`
	InterventionName string       `gorm:"not null" json:"intervention_name,omitempty"`
	ODKs             []ODK        `gorm:"many2many:consultation_odks;" json:"odks,omitempty"`
	Specialists      []Specialist `gorm:"many2many:consultation_specialists;" json:"specialists,omitempty"`
}

type ConsultationODK struct {
	ConsultationID uint `gorm:"primaryKey" json:"consultation_id,omitempty"`
	ODKID          uint `gorm:"primaryKey" json:"odk_id,omitempty"`
}

type ConsultationSpecialist struct {
	ConsultationID uint `gorm:"primaryKey" json:"consultation_id,omitempty"`
	SpecialistID   uint `gorm:"primaryKey" json:"specialist_id,omitempty"`
}

type Request struct {
	ClassNumber      int    `json:"classNumber,omitempty"`
	Class            string `json:"class,omitempty"`
	Code             string `json:"code,omitempty"`
	InterventionName string `json:"interventionName,omitempty"`
	ODKList          string `json:"ODKList,omitempty"`
	SpecialistList   string `json:"specialistList,omitempty"`
}

type LabTest struct {
	ID    uint   `gorm:"primaryKey" json:"id,omitempty"`
	Code  string `gorm:"not null" json:"code,omitempty"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Rule  string `gorm:"not null" json:"rule,omitempty"`
	Class string `gorm:"not null" json:"class,omitempty"`
}

type InstrumentalDiagnosticRequest struct {
	ClassNumber      int    `json:"classNumber,omitempty"`
	Class            string `json:"class,omitempty"`
	Code             string `json:"code,omitempty"`
	InterventionName string `json:"interventionName,omitempty"`
	SpecialistList   string `json:"specialistList,omitempty"`
}

type InstrumentalDiagnostic struct {
	ID          uint         `gorm:"primaryKey"`
	ClassNumber int          `gorm:"not null" json:"class_number,omitempty"`
	Name        string       `gorm:"not null" json:"name,omitempty"`
	Code        string       `gorm:"unique;not null" json:"code,omitempty"`
	FullName    string       `gorm:"not null" json:"full_name,omitempty"`
	Specialists []Specialist `gorm:"many2many:instrumental_diagnostic_specialists;" json:"specialists,omitempty"`
}
