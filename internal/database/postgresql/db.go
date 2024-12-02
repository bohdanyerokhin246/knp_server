package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"knp_server/internal/config"
	"os"
)

var DBTest *gorm.DB
var DBSite *gorm.DB
var DBUser *gorm.DB
var DBPost *gorm.DB
var DBStatistic *gorm.DB
var DBMedical *gorm.DB
var DBStorage *gorm.DB

func Connect() {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("hostDB"),
		os.Getenv("portDB"),
		os.Getenv("userDB"),
		os.Getenv("passwordDB"),
		os.Getenv("nameDB"),
		os.Getenv("sslModeDB"))

	DBSite, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "site."},
	})

	DBUser, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "user."},
	})

	DBPost, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "post."},
	})

	DBMedical, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "flg."},
	})

	DBStatistic, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "statistic."},
	})

	DBStorage, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "storage."},
	})

	err = DBSite.AutoMigrate(
		&config.Menu{},
		&config.Page{})
	if err != nil {
		fmt.Printf("Error with migration DBPost. Error: %v", err.Error())
	}

	err = DBUser.AutoMigrate(&config.User{})
	if err != nil {
		fmt.Printf("Error with migration DBPost. Error: %v", err.Error())
	}

	err = DBPost.AutoMigrate(&config.Post{})
	if err != nil {
		fmt.Printf("Error with migration DBPost. Error: %v", err.Error())
	}

	err = DBStatistic.AutoMigrate(
		&config.EMZ{},
		&config.StatisticPatient{})
	if err != nil {
		fmt.Printf("Error with migration DBStatistic. Error: %v", err.Error())
	}

	err = DBMedical.AutoMigrate(
		&config.Patient{},
		&config.Therapist{},
		&config.Diagnose{},
		&config.Exam{})
	if err != nil {
		fmt.Printf("Error with migration DBMedical. Error: %v", err.Error())
	}

	err = DBStorage.AutoMigrate(
		&config.Monitor{},
		&config.Cartridge{},
		&config.Computer{},
		&config.CabinetCard{},
		&config.StorageDevice{},
		&config.Processor{},
		&config.Contract{},
		&config.Periphery{},
		&config.Movement{},
		&config.Printer{},
		&config.Repair{},
		&config.RespPerson{},
		&config.RAM{})
	if err != nil {
		fmt.Printf("Error with migration DBSorage. Error: %v", err.Error())
	}
}
