package queries

import (
	"fmt"
	"gorm.io/gorm"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
	"log"
	"strings"
)

func FetchData[T any](result *[]T, conditions ...func(db *gorm.DB) *gorm.DB) error {
	db := postgresql.DB.MedInfo
	for _, condition := range conditions {
		db = condition(db)
	}
	return db.Find(result).Error
}

func FetchDataByID[T any](result *T, conditions ...func(db *gorm.DB) *gorm.DB) *gorm.DB {
	db := postgresql.DB.MedInfo
	for _, condition := range conditions {
		db = condition(db)
	}
	return db.Find(result)
}

//ODK Classes

func GetODKClasses() ([]models.ODK, error) {
	var odks []models.ODK
	if err := FetchData(&odks); err != nil {
		return nil, err
	}
	return odks, nil
}

func GetODKDiagnoses() ([]models.ODK, error) {
	var odks []models.ODK
	if err := FetchData(&odks, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Diagnoses")
	}); err != nil {
		return nil, err
	}
	return odks, nil
}

func GetODKClassById(id uint) (*models.ODK, error) {
	var odk models.ODK
	if err := postgresql.DB.MedInfo.First(&odk, id).Error; err != nil {
		return nil, err
	}
	return &odk, nil
}

//DiagnosesByODK

func GetODKDiagnoseById(id uint) (*models.ODK, error) {
	var odk models.ODK
	if err := postgresql.DB.MedInfo.Preload("Diagnoses").First(&odk, id).Error; err != nil {
		return nil, err
	}
	return &odk, nil
}

//Specialists

func GetSpecialists() ([]models.Specialist, error) {
	var specialists []models.Specialist

	resp := postgresql.DB.MedInfo.Where("status = ?", true).Find(&specialists)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return specialists, nil
}

func GetSpecialistByCode(id string) (*models.Specialist, error) {
	var specialist *models.Specialist

	resp := postgresql.DB.MedInfo.Where("code = ?", id).Find(&specialist)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return specialist, nil
}

//Lab services

func GetLabTests() ([]models.LabTest, error) {
	var labTests []models.LabTest

	resp := postgresql.DB.MedInfo.Find(&labTests)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return labTests, nil
}

func GetLabTestByCode(labTestCode string) (*models.LabTest, error) {
	var labTest *models.LabTest

	resp := postgresql.DB.MedInfo.Where("code = ?", labTestCode).Find(&labTest)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return labTest, nil
}

//Procedures

func GetProcedures() ([]models.Procedure, error) {
	var procedures []models.Procedure
	if err := postgresql.DB.MedInfo.Preload("ODKs").Preload("Specialists", "status = true").Order("class_number asc").Find(&procedures).Error; err != nil {
		return nil, err
	}
	return procedures, nil
}

func GetProceduresByCode(code string) ([]models.Procedure, error) {
	var procedures []models.Procedure

	if err := postgresql.DB.MedInfo.Preload("ODKs").Preload("Specialists", "status = true").Order("class_number asc").Where("code = ?", code).Find(&procedures).Error; err != nil {
		return nil, err
	}
	return procedures, nil
}
func GetProceduresBySpecialistCode(specialistCode string) ([]models.Procedure, error) {
	var procedures []models.Procedure
	if err := postgresql.DB.MedInfo.
		Joins("JOIN medical.procedure_specialists ps ON ps.procedure_id = procedures.id").
		Joins("JOIN medical.specialists s ON s.id = ps.specialist_id").
		Where("s.code = ? OR s.code = 'P9999'", specialistCode).
		Preload("ODKs").
		Preload("Specialists", "status = true AND code = ?", specialistCode).
		Order("medical.procedures.class_number, medical.procedures.intervention_name asc").
		Find(&procedures).Error; err != nil {
		return nil, err
	}
	return procedures, nil
}

func CreateProcedure(data []models.Request) ([]uint, []uint, []uint) {
	var proceduresID []uint
	var procedureODKID []uint
	var procedureSpecialistID []uint

	for _, d := range data {
		var odkIDs []uint
		var specialistIDs []uint

		procedure := models.Procedure{
			ClassNumber:      d.ClassNumber,
			Class:            d.Class,
			Code:             d.Code,
			InterventionName: d.InterventionName,
		}
		err := postgresql.DB.MedInfo.Create(&procedure)
		if err.Error != nil {
			log.Fatal("Procedure insertion error:", err)
		}
		proceduresID = append(proceduresID, procedure.ID)

		odkNumbers := strings.Split(d.ODKList, ",")
		for _, num := range odkNumbers {
			var odk models.ODK
			if err := postgresql.DB.MedInfo.Where("id = ?", num).First(&odk).Error; err == nil {
				odkIDs = append(odkIDs, odk.ID)
			}
		}

		for _, odkID := range odkIDs {
			procedureODK := models.ProcedureODK{ProcedureID: procedure.ID, ODKID: odkID}
			postgresql.DB.MedInfo.FirstOrCreate(&procedureODK, procedureODK)
			procedureODKID = append(procedureODKID, procedureODK.ProcedureID)
		}

		specialistCodes := strings.Split(d.SpecialistList, ",")
		for _, code := range specialistCodes {
			var specialist models.Specialist
			if err := postgresql.DB.MedInfo.Where("code = ?", code).First(&specialist).Error; err == nil {
				specialistIDs = append(specialistIDs, specialist.ID)
			}
		}

		for _, specialistID := range specialistIDs {
			procedureSpecialist := models.ProcedureSpecialist{ProcedureID: procedure.ID, SpecialistID: specialistID}
			postgresql.DB.MedInfo.FirstOrCreate(&procedureSpecialist, procedureSpecialist)
			procedureSpecialistID = append(procedureSpecialistID, procedureSpecialist.SpecialistID)
		}
	}

	return proceduresID, procedureODKID, procedureSpecialistID
}

//consultation

func GetConsultations() ([]models.Consultation, error) {
	var consultations []models.Consultation
	if err := postgresql.DB.MedInfo.
		Preload("ODKs").
		Preload("Specialists", "status = true").
		Joins("JOIN medical.consultation_specialists ON consultation_specialists.consultation_id = consultations.id").
		Joins("JOIN medical.specialists ON specialists.id = consultation_specialists.specialist_id AND specialists.status = true").
		Order("class_number ASC").
		Find(&consultations).Error; err != nil {
		return nil, err
	}
	return consultations, nil
}

func GetConsultationsByCode(code string) ([]models.Consultation, error) {
	var consultations []models.Consultation

	if err := postgresql.DB.MedInfo.
		Preload("ODKs").
		Preload("Specialists", "status = true").
		Order("class_number asc").
		Where("code = ?", code).
		Find(&consultations).
		Error; err != nil {
		return nil, err
	}
	return consultations, nil
}
func GetConsultationsBySpecialistCode(specialistCode string) ([]models.Consultation, error) {
	var consultations []models.Consultation
	if err := postgresql.DB.MedInfo.
		Joins("JOIN medical.procedure_specialists ps ON ps.procedure_id = procedures.id").
		Joins("JOIN medical.specialists s ON s.id = ps.specialist_id").
		Where("s.code = ? OR s.code = 'P9999'", specialistCode).
		Preload("ODKs").
		Preload("Specialists", "status = true AND code = ?", specialistCode).
		Order("medical.procedures.class_number, medical.procedures.intervention_name asc").
		Find(&consultations).Error; err != nil {
		return nil, err
	}
	return consultations, nil
}

func CreateConsultation(data []models.Request) ([]uint, []uint, []uint) {
	var consultationID []uint
	var consultationODKID []uint
	var consultationSpecialistID []uint

	for _, d := range data {
		var odkIDs []uint
		var specialistIDs []uint

		consultation := models.Consultation{
			ClassNumber:      d.ClassNumber,
			ClassName:        d.Class,
			Code:             d.Code,
			InterventionName: d.InterventionName,
		}
		err := postgresql.DB.MedInfo.Create(&consultation)
		if err.Error != nil {
			log.Fatal("Ошибка вставки процедуры:", err)
		}
		consultationID = append(consultationID, consultation.ID)

		odkNumbers := strings.Split(d.ODKList, ",")
		for _, num := range odkNumbers {
			var odk models.ODK
			if err := postgresql.DB.MedInfo.Where("id = ?", num).First(&odk).Error; err == nil {
				odkIDs = append(odkIDs, odk.ID)
			}
		}

		for _, odkID := range odkIDs {
			consultationODK := models.ConsultationODK{ConsultationID: consultation.ID, ODKID: odkID}
			postgresql.DB.MedInfo.FirstOrCreate(&consultationODK, consultationODK)
			consultationODKID = append(consultationODKID, consultationODK.ConsultationID)
		}

		specialistCodes := strings.Split(d.SpecialistList, ",")
		for _, code := range specialistCodes {
			var specialist models.Specialist
			if err := postgresql.DB.MedInfo.Where("code = ?", code).First(&specialist).Error; err == nil {
				specialistIDs = append(specialistIDs, specialist.ID)
			}
		}

		for _, specialistID := range specialistIDs {
			consultationSpecialist := models.ConsultationSpecialist{ConsultationID: consultation.ID, SpecialistID: specialistID}
			postgresql.DB.MedInfo.FirstOrCreate(&consultationSpecialist, consultationSpecialist)
			consultationSpecialistID = append(consultationSpecialistID, consultationSpecialist.SpecialistID)
		}
	}

	return consultationID, consultationODKID, consultationSpecialistID
}

//instrumentalDiagnostic

func GetInstrumentalDiagnostic() ([]models.InstrumentalDiagnostic, error) {
	var instrumentalDiagnostic []models.InstrumentalDiagnostic

	if err := postgresql.DB.MedInfo.Preload("Specialists", "status = true").Order("class_number asc").Find(&instrumentalDiagnostic).Error; err != nil {
		return nil, err
	}
	return instrumentalDiagnostic, nil
}

func GetInstrumentalDiagnosticsBySpecialistCode(specialistCode string) ([]models.InstrumentalDiagnostic, error) {
	var diagnostics []models.InstrumentalDiagnostic
	if err := postgresql.DB.MedInfo.
		Joins("JOIN medical.instrumental_diagnostic_specialists ids ON ids.instrumental_diagnostic_id = instrumental_diagnostics.id").
		Joins("JOIN medical.specialists s ON s.id = ids.specialist_id").
		Where("s.code = ?", specialistCode).
		Preload("Specialists", "status = true").
		Order("medical.instrumental_diagnostics.name ASC").
		Find(&diagnostics).Error; err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func CreateInstrumentalDiagnostic(data []models.InstrumentalDiagnosticRequest) ([]uint, []uint) {
	var diagnosticIDs []uint
	var specialistIDs []uint

	for _, d := range data {
		var specialists []models.Specialist

		specialistCodes := strings.Split(d.SpecialistList, ",")
		if err := postgresql.DB.MedInfo.Where("code IN ?", specialistCodes).Find(&specialists).Error; err != nil {
			log.Fatal("Ошибка при получении специалистов:", err)
		}

		if len(specialists) == 0 {
			log.Println("Предупреждение: не найдено специалистов для кода:", d.Code)
			continue
		}

		err := postgresql.DB.MedInfo.Transaction(func(tx *gorm.DB) error {
			diagnostic := models.InstrumentalDiagnostic{
				ClassNumber: d.ClassNumber,
				Name:        d.Class,
				Code:        d.Code,
				FullName:    d.InterventionName,
				Specialists: specialists,
			}

			if err := tx.Create(&diagnostic).Error; err != nil {
				return fmt.Errorf("ошибка вставки диагностики: %w", err)
			}

			diagnosticIDs = append(diagnosticIDs, diagnostic.ID)
			for _, s := range specialists {
				specialistIDs = append(specialistIDs, s.ID)
			}
			return nil
		})

		if err != nil {
			log.Fatal("Transaction error:", err)
		}
	}

	return diagnosticIDs, specialistIDs
}
