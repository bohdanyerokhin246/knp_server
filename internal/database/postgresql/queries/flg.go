package queries

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreatePatient(patients []models.Patient) error {

	postgresql.DB.Medical.Exec(`TRUNCATE flg.patients RESTART IDENTITY`)

	for _, patient := range patients {
		result := postgresql.DB.Medical.Create(&patient)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func CreateDiagnose(diagnoses []models.DiagnoseFLG) error {

	for _, diagnose := range diagnoses {
		result := postgresql.DB.Medical.Create(&diagnose)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func CreateExam(exams []models.Exam) error {

	postgresql.DB.Medical.Exec(`TRUNCATE flg.exams RESTART IDENTITY`)

	for _, exam := range exams {
		result := postgresql.DB.Medical.Create(&exam)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func CreateTherapist(therapists []models.Therapist) error {

	postgresql.DB.Medical.Exec(`TRUNCATE flg.therapists RESTART IDENTITY`)

	for _, therapist := range therapists {
		result := postgresql.DB.Medical.Create(&therapist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetExams() ([]models.ExamDetails, error) {

	var exams []models.ExamDetails

	err := postgresql.DB.Medical.Table("flg.exams").
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
