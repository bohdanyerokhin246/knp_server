package config

import (
	"gorm.io/gorm"
	"time"
)

type CabinetCard struct {
	gorm.Model
	RespPersonID uint `json:"respPersonsID,omitempty"`
	ComputerID   uint `json:"computerID,omitempty"`
	PrinterID    uint `json:"printerID,omitempty"`
	OtherID      uint `json:"othersID,omitempty"`
}

type Cartridge struct {
	gorm.Model
	Brand     string `json:"name,omitempty"`
	ModelName string `json:"modelName,omitempty"`
}

type Computer struct {
	gorm.Model
	IsChecked    string `json:"isChecked,omitempty"`
	Cabinet      string `json:"cabinet,omitempty"`
	WorkPlace    string `json:"workPlace,omitempty"`
	Brand        string `json:"brand,omitempty"`
	ModelName    string `json:"modelName,omitempty"`
	InvNumber    uint   `json:"invNumber,omitempty"`
	SerNumber    string `json:"serNumber,omitempty"`
	Power        uint   `json:"power,omitempty"`
	ProcessorID  uint   `json:"processorID,omitempty"`
	HddID        uint   `json:"hddID,omitempty"`
	SsdID        uint   `json:"ssdID,omitempty"`
	RamOneID     uint   `json:"ramOneID,omitempty"`
	RamTwoID     uint   `json:"ramTwoID,omitempty"`
	MonitorID    uint   `json:"monitorID,omitempty"`
	KeyboardID   uint   `json:"keyboardID,omitempty"`
	MouseID      uint   `json:"mouseID,omitempty"`
	ContractID   uint   `json:"contractID,omitempty"`
	RespPersonID uint   `json:"respPersonID,omitempty"`
	Quantity     uint   `json:"quantity,omitempty"`
}

type ComputerFullInfo struct {
	gorm.Model
	Cabinet        string `json:"cabinet,omitempty"`
	Brand          string `json:"brand,omitempty"`
	ModelName      string `json:"modelName,omitempty"`
	InvNumber      uint   `json:"invNumber,omitempty"`
	SerNumber      string `json:"serNumber,omitempty"`
	Power          uint   `json:"power,omitempty"`
	ProcessorName  string `json:"processorName,omitempty"`
	HddID          uint   `json:"hddID,omitempty"`
	SsdName        string `json:"ssdName,omitempty"`
	RamOneName     string `json:"ramOneName,omitempty"`
	RamTwoID       uint   `json:"ramTwoID,omitempty"`
	MonitorName    string `json:"monitorName,omitempty"`
	KeyboardName   string `json:"keyboardName,omitempty"`
	MouseID        uint   `json:"mouseID,omitempty"`
	ContractName   string `json:"contractName,omitempty"`
	RespPersonName string `json:"respPersonName,omitempty"`
	Quantity       uint   `json:"quantity,omitempty"`
}

type Contract struct {
	gorm.Model
	Name          string    `json:"name,omitempty"`
	Firm          string    `json:"firm,omitempty"`
	Date          time.Time `json:"date"`
	ServerAddress string    `json:"serverAddress,omitempty"`
}

type StorageDevice struct {
	gorm.Model
	Brand     string `json:"brand,omitempty"`
	ModelName string `json:"modelName,omitempty"`
	Capacity  int    `json:"capacity,omitempty"`
}

type RAM struct {
	gorm.Model
	Brand     string `json:"brand,omitempty"`
	ModelName string `json:"modelName,omitempty"`
	Capacity  int    `json:"capacity,omitempty"`
}

type Processor struct {
	gorm.Model
	Brand     string  `json:"brand,omitempty"`
	ModelName string  `json:"modelName,omitempty"`
	Frequency float32 `json:"frequency,omitempty"`
	Power     uint    `json:"power,omitempty"`
	Socket    string  `json:"socket,omitempty"`
}

type Periphery struct {
	gorm.Model
	Brand        string `json:"brand,omitempty"`
	ModelName    string `json:"modelName,omitempty"`
	InvNumber    uint   `json:"invNumber,omitempty"`
	ContractID   uint   `json:"contractID,omitempty"`
	RespPersonID uint   `json:"respPersonID,omitempty"`
}

type Monitor struct {
	gorm.Model
	Brand        string  `json:"brand,omitempty"`
	ModelName    string  `json:"modelName,omitempty"`
	InvNumber    uint    `json:"invNumber,omitempty"`
	SerNumber    string  `json:"serNumber,omitempty"`
	Size         float32 `json:"size,omitempty"`
	ContractID   uint    `json:"contractID,omitempty"`
	RespPersonID uint    `json:"respPersonID,omitempty"`
	Power        uint    `json:"power,omitempty"`
}

type Movement struct {
	gorm.Model
	ProductID  uint      `json:"productID,omitempty"`
	Quantity   uint      `json:"quantity,omitempty"`
	ForWhoID   uint      `json:"forWhoID,omitempty"`
	FromWhomID uint      `json:"fromWhomID,omitempty"`
	Date       time.Time `json:"date"`
}

type Printer struct {
	gorm.Model
	Brand        string `json:"brand,omitempty"`
	ModelName    string `json:"modelName,omitempty"`
	InvNumber    uint   `json:"invNumber,omitempty"`
	SerNumber    string `json:"serNumber,omitempty"`
	CartridgeID  uint   `json:"cartridgeID,omitempty"`
	ContractID   uint   `json:"contractID,omitempty"`
	RespPersonID uint   `json:"respPersonID,omitempty"`
}

type Repair struct {
	gorm.Model
	ProductInvNumber string `json:"productInvNumber,omitempty"`
	ProductSerNumber string `json:"productSerNumber,omitempty"`
	Reason           string `json:"reason,omitempty"`
	TransferDate     string `json:"transferDate"`
	ReceivingDate    string `json:"receivingDate"`
}
type RespPerson struct {
	gorm.Model
	Name       string `json:"name,omitempty"`
	JobTitle   string `json:"jobTitle,omitempty"`
	Department string `json:"department,omitempty"`
}
