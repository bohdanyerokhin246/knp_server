package postgresql

import (
	"knp_server/internal/config"
)

func CreateMonitor(monitor config.Monitor) error {

	result := DBStorage.Create(&monitor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateCabinetCard(cabinetCard config.CabinetCard) error {

	result := DBStorage.Create(&cabinetCard)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateCartridge(cartridge config.Cartridge) error {

	result := DBStorage.Create(&cartridge)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateStorageDevice(device config.StorageDevice) error {

	result := DBStorage.Create(&device)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateRAM(ram config.RAM) error {

	result := DBStorage.Create(&ram)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateComputer(computer config.Computer) error {

	result := DBStorage.Create(&computer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateContract(contract config.Contract) error {

	result := DBStorage.Create(&contract)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateMovement(movement config.Movement) error {

	result := DBStorage.Create(&movement)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreatePeriphery(periphery config.Periphery) error {

	result := DBStorage.Create(&periphery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreatePrinter(printer config.Printer) error {

	result := DBStorage.Create(&printer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateProcessor(processor config.Processor) error {

	result := DBStorage.Create(&processor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateResponsePerson(person config.RespPerson) error {

	result := DBStorage.Create(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetComputers() ([]config.Computer, error) {

	var computers []config.Computer

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DBStorage.Order("id asc").Find(&computers)

	if err.Error != nil {
		return nil, err.Error
	}

	return computers, err.Error
}

func UpdateComputers(computers []config.Computer) error {

	for _, computer := range computers {
		err := DBStorage.
			Model(&computer).
			Updates(config.Computer{
				IsChecked:   computer.IsChecked,
				Cabinet:     computer.Cabinet,
				WorkPlace:   computer.WorkPlace,
				InvNumber:   computer.InvNumber,
				SerNumber:   computer.SerNumber,
				Brand:       computer.Brand,
				ModelName:   computer.ModelName,
				ProcessorID: computer.ProcessorID,
				MonitorID:   computer.MonitorID,
				SsdID:       computer.SsdID,
			})

		if err.Error != nil {
			return err.Error
		}
	}

	return nil
}

func GetMonitors() ([]config.Monitor, error) {

	var monitors []config.Monitor

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DBStorage.Order("id asc").Find(&monitors)

	if err.Error != nil {
		return nil, err.Error
	}

	return monitors, err.Error
}

func ReadExams() ([]config.ExamDetails, error) {

	var exams []config.ExamDetails

	err := DBMedical.Table("exams").
		Limit(1000).
		Select("exams.exam_id, exams.exam_date, diagnoses.diagnose AS diagnose_name,patients.full_name AS full_name, therapists.full_name AS therapist_name").
		Joins("left join patients on patients.id = exams.patient_id").
		Joins("left join diagnoses on diagnoses.id = exams.diagnose_id").
		Joins("left join therapists on therapists.id = exams.therapist_id").
		Scan(&exams).Error

	if err != nil {
		return nil, err
	}

	return exams, nil
}

func GetComputerFullInfo() ([]config.ComputerFullInfo, error) {
	var computerFullInfo []config.ComputerFullInfo

	err := DBStorage.Table("storage.computers").
		Select("cabinet, brand, model_name,inv_number, power, processor_name").
		Joins("left join storage.processors on processors.id = computers.processor_id").
		Joins("left join storage.storage_devices on storage_devices.id = computers.ssd_id").
		Joins("left join storage.rams on rams.id = computers.ram_one_id").
		Joins("left join storage. monitors on monitors.id = computers.monitor_id").
		Joins("left join storage.peripheries on peripheries.id = computers.keyboard_id").
		Joins("left join storage.contracts on contracts.id = computers.contract_id").
		Joins("left join storage.resp_people on resp_people.id = computers.resp_person_id").
		Scan(&computerFullInfo).Error

	if err != nil {
		return nil, err
	}

	return computerFullInfo, nil
}

func CreateRepair(repair config.Repair) error {

	result := DBStorage.Create(&repair)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
