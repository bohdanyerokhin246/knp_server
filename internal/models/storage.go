package models

import "gorm.io/gorm"

type Subdivision struct {
	ID         uint     `json:"id,omitempty" gorm:"primaryKey"`
	Name       string   `json:"name,omitempty"`
	EmployeeID uint     `json:"employeeID,omitempty"`
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
}

type Employee struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type Cabinet struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"size:100;not null"`
}

type EquipmentEmployee struct {
	EquipmentID uint `gorm:"primaryKey" json:"specialist_id,omitempty"`
	EmployeeID  uint `gorm:"primaryKey" json:"employee_id,omitempty"`
}

type Equipment struct {
	gorm.Model
	Type            string   `json:"typeName,omitempty"`
	Name            string   `json:"name,omitempty"`
	SerialNumber    string   `json:"serialNumber,omitempty"`
	InventoryNumber string   `json:"inventoryNumber,omitempty"`
	Brand           string   `json:"brand,omitempty"`
	ModelName       string   `json:"modelName,omitempty"`
	Power           int      `json:"power,omitempty"`
	VA              int      `json:"va,omitempty"`
	CabinetID       uint     `json:"cabinetID,omitempty"`
	Cabinet         Cabinet  `gorm:"foreignKey:CabinetID"`
	EmployeeID      uint     `json:"employeeID,omitempty"`
	Employee        Employee `gorm:"foreignKey:EmployeeID"`
}

type Monitor struct {
	ID              uint      `json:"id,omitempty" gorm:"primaryKey"`
	EquipmentID     uint      `json:"equipmentID,omitempty" json:"equipmentID,omitempty"`
	Equipment       Equipment `gorm:"foreignKey:EquipmentID"`
	ScreenSize      int       `json:"screenSize,omitempty"`
	Resolution      string    `json:"resolution,omitempty" gorm:"size:50"`
	SerialNumber    string    `json:"serialNumber,omitempty" gorm:"size:50"`
	InventoryNumber string    `json:"inventoryNumber,omitempty" gorm:"size:50"`
	Brand           string    `json:"brand,omitempty" gorm:"size:100"`
	Model           string    `json:"model,omitempty" gorm:"size:100"`
	Power           int       `json:"power,omitempty"`
}

type Computer struct {
	ID            uint            `json:"id,omitempty" gorm:"primaryKey"`
	EquipmentID   uint            `json:"equipmentID,omitempty"`
	Equipment     Equipment       `gorm:"foreignKey:EquipmentID"`
	ProcessorID   uint            `json:"processorID,omitempty"`
	Processor     Processor       `gorm:"foreignKey:ProcessorID"`
	RAMModules    []RAMModule     `gorm:"many2many:computer_ram_modules;"`
	StorageDrives []StorageDevice `gorm:"many2many:computer_storage_drives;"`
	Power         int             `json:"power,omitempty"`
}

type Processor struct {
	ID        uint   `json:"id,omitempty" gorm:"primaryKey"`
	Brand     string `json:"brand,omitempty" gorm:"size:100"`
	Model     string `json:"model,omitempty" gorm:"size:100"`
	Cores     int    `json:"cores,omitempty"`
	BaseClock string `json:"baseClock,omitempty" gorm:"size:50"`
}

type RAMModule struct {
	ID    uint   `json:"id,omitempty" gorm:"primaryKey"`
	Brand string `json:"brand,omitempty" gorm:"size:100"`
	Model string `json:"model,omitempty" gorm:"size:100"`
	Size  string `json:"size,omitempty" gorm:"size:50"`
	Type  string `json:"type,omitempty" gorm:"size:50"`
}

type StorageDevice struct {
	ID       uint   `json:"id,omitempty" gorm:"primaryKey"`
	Brand    string `json:"brand,omitempty" gorm:"size:100"`
	Model    string `json:"model,omitempty" gorm:"size:100"`
	Type     string `json:"type,omitempty" gorm:"size:50"`
	Capacity string `json:"capacity,omitempty" gorm:"size:50"`
}

type ComputerRAM struct {
	ComputerID  uint
	RAMModuleID uint
}

type ComputerStorage struct {
	ComputerID     uint
	StorageDriveID uint
}

////////

//type Contract struct {
//	gorm.Model
//	Name          string    `json:"name,omitempty"`
//	Firm          string    `json:"firm,omitempty"`
//	Date          time.Time `json:"date"`
//	ServerAddress string    `json:"serverAddress,omitempty"`
//}
//
//type Movement struct {
//	gorm.Model
//	ProductID  uint      `json:"productID,omitempty"`
//	Quantity   uint      `json:"quantity,omitempty"`
//	ForWhoID   uint      `json:"forWhoID,omitempty"`
//	FromWhomID uint      `json:"fromWhomID,omitempty"`
//	Date       time.Time `json:"date"`
//}
//
//type Repair struct {
//	gorm.Model
//	RepairDate time.Time `json:"repairDate"`
//	Name       string    `json:"name,omitempty"`
//	Reason     string    `json:"reason,omitempty"`
//	Type       string    `json:"type,omitempty"`
//	InvNumber  string    `json:"invNumber,omitempty"`
//	SerNumber  string    `json:"serNumber,omitempty"`
//	Cabinet    string    `json:"cabinet,omitempty"`
//}
//
//type Cartridge struct {
//	gorm.Model
//	Brand     string `json:"name,omitempty"`
//	ModelName string `json:"modelName,omitempty"`
//}
