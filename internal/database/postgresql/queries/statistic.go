package queries

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm/clause"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
	"log"
)

func CreateStatisticPatient(patients []models.StatisticPatient) error {

	for _, patient := range patients {
		result := postgresql.DB.Statistic.Create(&patient)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func GetDepartmentByDoctor() (map[string]string, error) {
	rows, err := postgresql.DB.Statistic.Raw(
		`SELECT 
    			statistic.doctor_department.doctor_name,
    			statistic.doctor_department.department_name 
			FROM 
    			statistic.doctor_department
    		`).Rows()

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)

	result := make(map[string]string)

	var column1, column2 string
	for rows.Next() {
		err = rows.Scan(&column1, &column2)
		if err != nil {
			return nil, err
		}
		result[column1] = column2
	}

	return result, nil
}

func CreateEmzs(statistics []models.EMZ) error {
	for _, statistic := range statistics {
		//If emz_id unique create emz in statistic.emz
		//else ignore emz
		result := postgresql.DB.Statistic.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "emz_id"}},
			DoNothing: true,
		}).
			Create(&statistic)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func CorrectionEMZPaymentActual() error {

	var err error
	var patientsList []models.StatisticPatient
	var emzList []*models.EMZ

	//Get list of all patients from statistic.statistic_patients
	patientsList, err = getPatientsList()
	if err != nil {
		return err
	}

	//For each patient from patientsList get list of EMZ
	for _, patient := range patientsList {
		emzList, err = getAllEMZByPatientID(
			patient.Year,
			patient.Month,
			patient.PatientID,
			patient.Package)

		if err != nil {

			return err
		}

		//Patient`s capitation rate / amount of EMZ in this month = payment actual for EMZ
		emzList, err = divisionPaymentActualOfEMZ(emzList)

		if err != nil {
			return err
		}

		//Update in DB list of EMZ with price after division
		err = updatePaymentActualByEMZ(emzList)

		if err != nil {
			return err
		}
	}
	return nil
}

func getPatientsList() ([]models.StatisticPatient, error) {

	var patientsList []models.StatisticPatient

	err := postgresql.DB.Statistic.Order("patient_id asc").Find(&patientsList)

	if err.Error != nil {
		return nil, err.Error
	}

	return patientsList, nil
}

func getAllEMZByPatientID(year, month int, patientID, packageName string) ([]*models.EMZ, error) {

	var EMZList []*models.EMZ

	err := postgresql.DB.Statistic.
		Unscoped().
		Where(
			"year = ? AND month = ? AND patients_id = ? AND package = ?",
			year,
			month,
			patientID,
			packageName,
		).
		Find(&EMZList)

	if err.Error != nil {
		return nil, err.Error
	}

	return EMZList, nil
}

// DivisionPaymentActualOfEMZ get array of EMZ count len of this array,
// then changes payment_actual as follows:
// payment_actual = capitation_rate_for_package / len([]models.EMZ)
func divisionPaymentActualOfEMZ(EMZList []*models.EMZ) ([]*models.EMZ, error) {

	for _, emz := range EMZList {
		cost, err := getTariffByPatientID(emz.Month, emz.PatientsID, emz.Package)
		if err != nil {
			fmt.Println(emz.EmzID)
			return nil, err

		}

		if cost == 0 {
			emz.PaymentActual = cost
		} else {
			emz.PaymentActual = cost / float64(len(EMZList))
		}
	}

	return EMZList, nil
}

func getTariffByPatientID(month int, patientID, packageName string) (float64, error) {

	var patient models.StatisticPatient

	err := postgresql.DB.Statistic.Where("month = ? AND patient_id = ? AND package = ?", month, patientID, packageName).Find(&patient)

	if err.Error != nil {
		return 0.0, err.Error
	}

	return patient.PaymentActual, nil
}

func updatePaymentActualByEMZ(EMZList []*models.EMZ) error {

	for _, emz := range EMZList {
		err := postgresql.DB.Statistic.Save(emz)
		if err.Error != nil {
			fmt.Println(emz.EmzID)
			return err.Error
		}
	}
	return nil
}

//Statistic getters

func GetStatisticAll() ([]models.EMZ, error) {

	var statistics []models.EMZ

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := postgresql.DB.Statistic.Find(&statistics)

	if err.Error != nil {
		return nil, err.Error
	}

	return statistics, err.Error
}

func GetIncludedSummarizeStatistic() []models.SummarizeStatistic {
	var summarizeStatistic []models.SummarizeStatistic

	err := postgresql.DB.Statistic.Table("statistic.emzs").
		Select(
			`month,
					year, 
					doctor_name, 
					doctor_unit, 
					package, 
					COUNT(id) as count_patients_id, 
					COUNT(DISTINCT patients_id) as unique_patients, 
					SUM(amount_of_included_services) as sum_included, 
					AVG(amount_of_included_services) as avg_included, 
					SUM(cost_estimated) as cost_estimated, 
					SUM(payment_actual) as payment_actual`).
		Where("package != ?", "-").
		//Where("emz_included_to_statistic", "Так").
		//Where("deleted_at IS NULL").
		Group("month, year, doctor_name, doctor_unit, package").
		Order("doctor_name, package").
		Scan(&summarizeStatistic).
		Error

	if err != nil {
		log.Fatal("failed to execute query:", err)
	}

	return summarizeStatistic
}

func GetStatisticByUnit() ([]models.SummarizeStatistic, error) {
	var summarizeStatistic []models.SummarizeStatistic

	err := postgresql.DB.Statistic.Table("statistic.emzs").
		Select(`
					year, 
					month,
					doctor_unit,
					package,
					COUNT(id) as count_patients_id, 
					COUNT(DISTINCT patients_id) as unique_patients, 
					SUM(amount_of_included_services) as sum_included, 
					AVG(amount_of_included_services) as avg_included, 
					SUM(cost_estimated) as cost_estimated, 
					SUM(payment_actual) as payment_actual`).
		//Where("package != ?", "-").
		//Where("emz_included_to_statistic", "Так").
		//Where("deleted_at IS NULL").
		Group("year, month, doctor_unit, package").
		//Order("doctor_unit, package").
		Scan(&summarizeStatistic).
		Error

	if err != nil {
		return nil, err
	}

	return summarizeStatistic, nil
}

func GetStatisticByPackage() ([]models.EMZ, error) {
	var summarizedEMZ []models.EMZ

	err := postgresql.DB.Statistic.Model(&models.EMZ{}).
		Select("year, month, package, SUM(payment_actual) as payment_actual").
		Where("package != ?", "-").
		Group("year, month, package").
		Scan(&summarizedEMZ).Error

	if err != nil {
		return nil, err
	}

	return summarizedEMZ, nil
}
