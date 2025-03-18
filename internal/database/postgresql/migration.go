package postgresql

import (
	"fmt"
	"knp_server/internal/models"
)

func Migrate() error {
	var err error

	err = DB.Site.AutoMigrate(
		&models.Menu{},
		&models.Page{},
		&models.MenuItem{})
	if err != nil {
		return fmt.Errorf("Error with migration Site: %v\n", err)
	}

	err = DB.User.AutoMigrate(
		&models.User{})
	if err != nil {
		return fmt.Errorf("Error with migration User: %v\n", err)
	}

	err = DB.Post.AutoMigrate(
		&models.News{},
		&models.FileInfo{})
	if err != nil {
		return fmt.Errorf("Error with migration News: %v\n", err)
	}

	err = DB.Statistic.AutoMigrate(
		&models.EMZ{},
		&models.StatisticPatient{})
	if err != nil {
		return fmt.Errorf("Error with migration Statistic: %v\n", err)
	}

	err = DB.Medical.AutoMigrate(
		&models.Patient{},
		&models.Therapist{},
		&models.DiagnoseFLG{},
		&models.Exam{})
	if err != nil {
		return fmt.Errorf("Error with migration Medical: %v\n", err)
	}

	err = DB.Storage.AutoMigrate(
		&models.Subdivision{},
		&models.Employee{},
		&models.Cabinet{},
		&models.Equipment{},
		&models.Monitor{},
		&models.Processor{},
		&models.RAMModule{},
		&models.ComputerRAM{},
		&models.StorageDevice{},
		&models.ComputerStorage{},
	)
	if err != nil {
		return fmt.Errorf("Error with migration Storage: %v\n", err)
	}

	err = DB.MedInfo.AutoMigrate(
		&models.Procedure{},
		&models.ProcedureODK{},
		&models.ProcedureSpecialist{},
		&models.InstrumentalDiagnostic{},
		&models.Consultation{},
		&models.ConsultationODK{},
		&models.ConsultationSpecialist{},
		&models.InstrumentalDiagnostic{})
	if err != nil {
		return fmt.Errorf("Error with migration MedInfo: %v\n", err)
	}

	return nil
}
