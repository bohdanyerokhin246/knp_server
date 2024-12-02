package postgresql

import (
	"knp_server/internal/config"
)

func CreatePatient(patients []config.Patient) error {

	DBMedical.Exec(`TRUNCATE flg.patients RESTART IDENTITY`)

	for _, patient := range patients {
		result := DBMedical.Create(&patient)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func CreateDiagnose(diagnoses []config.Diagnose) error {

	for _, diagnose := range diagnoses {
		result := DBMedical.Create(&diagnose)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func CreateExam(exams []config.Exam) error {

	DBMedical.Exec(`TRUNCATE flg.exams RESTART IDENTITY`)

	for _, exam := range exams {
		result := DBMedical.Create(&exam)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func CreateTherapist(therapists []config.Therapist) error {

	DBMedical.Exec(`TRUNCATE flg.therapists RESTART IDENTITY`)

	for _, therapist := range therapists {
		result := DBMedical.Create(&therapist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetExams() ([]config.ExamDetails, error) {

	var exams []config.ExamDetails

	err := DBMedical.Table("flg.exams").
		Limit(1000).
		Select("exams.exam_id, exams.exam_date, flg.diagnoses.diagnose AS diagnose_name, flg.patients.full_name AS full_name, flg.therapists.full_name AS therapist_name").
		Joins("left join flg.patients on flg.patients.id = flg.exams.patient_id").
		Joins("left join flg.diagnoses on flg.diagnoses.id = flg.exams.diagnose_id").
		Joins("left join flg.therapists on flg.therapists.id = flg.exams.therapist_id").
		Scan(&exams).Error

	if err != nil {
		return nil, err
	}

	return exams, nil
}
