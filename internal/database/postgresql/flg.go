package postgresql

import (
	"knp_server/internal/config"
)

func CheckPatientUnique(patient config.Patient) error {
	result := DB.Where("lastname = ?", patient.Lastname).First(&patient)
	if result.Error != nil {
		_ = CreatePatient(patient)
	}

	return nil
}

func CreatePatient(patient config.Patient) error {

	result := DB.Create(&patient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
