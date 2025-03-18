package storage

import (
	"errors"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreateComputer(computer *models.Computer) error {

	tx := postgresql.DB.Storage.Begin()

	var equipment models.Equipment
	if err := tx.First(&equipment, computer.EquipmentID).Error; err != nil {
		tx.Rollback()
		return errors.New("the specified equipment was not found")
	}
	computer.Equipment = equipment

	var processor models.Processor
	if err := tx.First(&processor, computer.ProcessorID).Error; err != nil {
		tx.Rollback()
		return errors.New("the specified processor was not found")
	}
	computer.Processor = processor

	if err := tx.Create(computer).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(computer.RAMModules) > 0 {
		var ramModules []models.RAMModule
		for _, ram := range computer.RAMModules {
			var existingRAM models.RAMModule
			if err := tx.First(&existingRAM, ram.ID).Error; err != nil {
				tx.Rollback()
				return errors.New("one of the specified RAM modules was not found")
			}
			ramModules = append(ramModules, existingRAM)
		}
		if err := tx.Model(computer).Association("RAMModules").Append(ramModules); err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(computer.StorageDrives) > 0 {
		var storageDrives []models.StorageDevice
		for _, storage := range computer.StorageDrives {
			var existingStorage models.StorageDevice
			if err := tx.First(&existingStorage, storage.ID).Error; err != nil {
				tx.Rollback()
				return errors.New("one of the specified drives was not found\"")
			}
			storageDrives = append(storageDrives, existingStorage)
		}
		if err := tx.Model(computer).Association("StorageDrives").Append(storageDrives); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func GetComputerByID(computerID uint) (*models.Computer, error) {
	var computer models.Computer

	err := postgresql.DB.Storage.Preload("Equipment").
		Preload("Equipment.Cabinet").
		Preload("Processor").
		Preload("RAMModules").
		Preload("StorageDrives").
		First(&computer, computerID).Error

	if err != nil {
		return nil, err
	}

	return &computer, nil
}

//
//import (
//	"knp_server/internal/models"
//)
//
//func CreateNomenclature(nomenclature models.Nomenclature) error {
//
//	result := DB.Storage.Create(&nomenclature)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func GetNomenclatures() ([]models.Nomenclature, error) {
//
//	var nomenclatures []models.Nomenclature
//
//	err := DB.Storage.Find(&nomenclatures)
//	if err.Error != nil {
//		return nil, err.Error
//	}
//	return nomenclatures, nil
//}
//
//func UpdateNomenclature(nomenclature models.Nomenclature) error {
//
//	result := DB.Storage.Save(&nomenclature)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}
//
//func DeleteNomenclature(nomenclature models.Nomenclature) error {
//
//	result := DB.Storage.Delete(&nomenclature)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateMonitor(monitor models.Monitor) error {
//
//	result := DB.Storage.Create(&monitor)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func GetMonitors() ([]models.Monitor, error) {
//
//	var monitors []models.Monitor
//
//	err := DB.Storage.Find(&monitors)
//	if err.Error != nil {
//		return nil, err.Error
//	}
//	return monitors, nil
//}
//
//func UpdateMonitor(monitors []models.Monitor) error {
//
//	for _, monitor := range monitors {
//		result := DB.Storage.Save(&monitor)
//		if result.Error != nil {
//			return result.Error
//		}
//	}
//	return nil
//}
//
//func DeleteMonitor(monitor models.Monitor) error {
//
//	result := DB.Storage.Create(&monitor)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateCabinetCard(cabinetCard models.CabinetCard) error {
//
//	result := DB.Storage.Create(&cabinetCard)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateCartridge(cartridge models.Cartridge) error {
//
//	result := DB.Storage.Create(&cartridge)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateStorageDevice(device models.StorageDevice) error {
//
//	result := DB.Storage.Create(&device)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateRAM(ram models.RAM) error {
//
//	result := DB.Storage.Create(&ram)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateComputer(computer models.Computer) error {
//
//	result := DB.Storage.Create(&computer)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateContract(contract models.Contract) error {
//
//	result := DB.Storage.Create(&contract)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateMovement(movement models.Movement) error {
//
//	result := DB.Storage.Create(&movement)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreatePeriphery(periphery models.Periphery) error {
//
//	result := DB.Storage.Create(&periphery)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreatePrinter(printer models.Printer) error {
//
//	result := DB.Storage.Create(&printer)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateProcessor(processor models.Processor) error {
//
//	result := DB.Storage.Create(&processor)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func CreateResponsePerson(person models.RespPerson) error {
//
//	result := DB.Storage.Create(&person)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func GetComputers() ([]models.Computer, error) {
//
//	var computers []models.Computer
//
//	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
//	err := DB.Storage.Order("id asc").Find(&computers)
//
//	if err.Error != nil {
//		return nil, err.Error
//	}
//
//	return computers, err.Error
//}
//
//func UpdateComputers(computers []models.Computer) error {
//
//	for _, computer := range computers {
//		err := DB.Storage.
//			Model(&computer).
//			Updates(models.Computer{
//				IsChecked:   computer.IsChecked,
//				Cabinet:     computer.Cabinet,
//				WorkPlace:   computer.WorkPlace,
//				InvNumber:   computer.InvNumber,
//				SerNumber:   computer.SerNumber,
//				Brand:       computer.Brand,
//				ModelName:   computer.ModelName,
//				ProcessorID: computer.ProcessorID,
//				MonitorID:   computer.MonitorID,
//				SsdID:       computer.SsdID,
//			})
//
//		if err.Error != nil {
//			return err.Error
//		}
//	}
//
//	return nil
//}
//
//func ReadExams() ([]models.ExamDetails, error) {
//
//	var exams []models.ExamDetails
//
//	err := DB.Medical.Table("exams").
//		Limit(1000).
//		Select("exams.exam_id, exams.exam_date, diagnoses.diagnose AS diagnose_name,patients.full_name AS full_name, therapists.full_name AS therapist_name").
//		Joins("left join patients on patients.id = exams.patient_id").
//		Joins("left join diagnoses on diagnoses.id = exams.diagnose_id").
//		Joins("left join therapists on therapists.id = exams.therapist_id").
//		Scan(&exams).Error
//
//	if err != nil {
//		return nil, err
//	}
//
//	return exams, nil
//}
//
//func GetComputerFullInfo() ([]models.ComputerFullInfo, error) {
//	var computerFullInfo []models.ComputerFullInfo
//
//	err := DB.Storage.Table("storage.computers").
//		Select("cabinet, brand, model_name,inv_number, power, processor_name").
//		Joins("left join storage.processors on processors.id = computers.processor_id").
//		Joins("left join storage.storage_devices on storage_devices.id = computers.ssd_id").
//		Joins("left join storage.rams on rams.id = computers.ram_one_id").
//		Joins("left join storage. monitors on monitors.id = computers.monitor_id").
//		Joins("left join storage.peripheries on peripheries.id = computers.keyboard_id").
//		Joins("left join storage.contracts on contracts.id = computers.contract_id").
//		Joins("left join storage.resp_people on resp_people.id = computers.resp_person_id").
//		Scan(&computerFullInfo).Error
//
//	if err != nil {
//		return nil, err
//	}
//
//	return computerFullInfo, nil
//}
//
///////
//
//func CreateRepair(repair models.Repair) error {
//
//	result := DB.Storage.Create(&repair)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func GetRepairs() ([]models.Repair, error) {
//
//	var repairs []models.Repair
//
//	err := DB.Storage.Find(&repairs)
//	if err.Error != nil {
//		return nil, err.Error
//	}
//	return repairs, nil
//}
//
//func UpdateRepair(repair models.Repair) error {
//
//	result := DB.Storage.Save(&repair)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}
//
//func DeleteRepair(repair models.Repair) error {
//
//	result := DB.Storage.Delete(&repair)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
