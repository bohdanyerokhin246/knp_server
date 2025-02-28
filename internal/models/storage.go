package models

type Cabinet struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"size:100;not null"`
}

type Equipment struct {
	ID              uint   `json:"id,omitempty" gorm:"primaryKey"`
	Type            string `json:"type,omitempty"`
	Name            string `json:"name,omitempty"`
	SerialNumber    string `json:"serialNumber,omitempty" gorm:"size:50"`
	InventoryNumber string `json:"inventoryNumber,omitempty" gorm:"size:50"`
	Brand           string `json:"brand,omitempty" gorm:"size:100"`
	Model           string `json:"model,omitempty" gorm:"size:100"`
	CabinetID       uint
	Cabinet         Cabinet `json:"cabinet,omitempty" gorm:"foreignKey:CabinetID"`
}

type Computer struct {
	ID          uint `json:"id,omitempty" gorm:"primaryKey"`
	EquipmentID uint
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`

	ProcessorID uint
	Processor   Processor `gorm:"foreignKey:ProcessorID"`

	RAMModules    []RAMModule     `gorm:"many2many:computer_ram_modules;"`
	StorageDrives []StorageDevice `gorm:"many2many:computer_storage_drives;"`
}

type Processor struct {
	ID        uint   `gorm:"primaryKey"`
	Brand     string `gorm:"size:100"`
	Model     string `gorm:"size:100"`
	Cores     int
	BaseClock string `gorm:"size:50"`
}

type RAMModule struct {
	ID    uint   `gorm:"primaryKey"`
	Brand string `gorm:"size:100"`
	Model string `gorm:"size:100"`
	Size  string `gorm:"size:50"` // Например, "8GB"
	Type  string `gorm:"size:50"` // Например, "DDR4"
	Speed string `gorm:"size:50"` // Например, "3200MHz"
}

type StorageDevice struct {
	ID       uint   `gorm:"primaryKey"`
	Brand    string `gorm:"size:100"`
	Model    string `gorm:"size:100"`
	Type     string `gorm:"size:50"` // HDD, SSD, NVMe
	Capacity string `gorm:"size:50"` // Например, "1TB"
}

type ComputerRAM struct {
	ComputerID  uint
	RAMModuleID uint
}

type ComputerStorage struct {
	ComputerID     uint
	StorageDriveID uint
}

type Peripheral struct {
	ID              uint `gorm:"primaryKey"`
	EquipmentID     uint
	Type            string    `gorm:"size:50"`
	SerialNumber    string    `gorm:"size:50"`
	InventoryNumber string    `gorm:"size:50"`
	Brand           string    `gorm:"size:100"`
	Model           string    `gorm:"size:100"`
	Equipment       Equipment `gorm:"foreignKey:EquipmentID"`
}

type Printer struct {
	ID              uint `gorm:"primaryKey"`
	EquipmentID     uint
	Type            string    `gorm:"size:50"`
	SerialNumber    string    `gorm:"size:50"`
	InventoryNumber string    `gorm:"size:50"`
	Brand           string    `gorm:"size:100"`
	Model           string    `gorm:"size:100"`
	Equipment       Equipment `gorm:"foreignKey:EquipmentID"`
}

type Monitor struct {
	ID              uint `gorm:"primaryKey"`
	EquipmentID     uint
	ScreenSize      int
	Resolution      string    `gorm:"size:50"`
	SerialNumber    string    `gorm:"size:50"`
	InventoryNumber string    `gorm:"size:50"`
	Brand           string    `gorm:"size:100"`
	Model           string    `gorm:"size:100"`
	Equipment       Equipment `gorm:"foreignKey:EquipmentID"`
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
//type RespPerson struct {
//	gorm.Model
//	Name       string `json:"name,omitempty"`
//	JobTitle   string `json:"jobTitle,omitempty"`
//	Department string `json:"department,omitempty"`
//}
//
//type Cartridge struct {
//	gorm.Model
//	Brand     string `json:"name,omitempty"`
//	ModelName string `json:"modelName,omitempty"`
//}
