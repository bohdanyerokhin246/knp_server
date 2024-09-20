package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"knp_server/internal/config"
	"os"
)

var DB *gorm.DB

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

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error conection to DB. Error: %v", err)
	}

	err = DB.AutoMigrate(&config.Post{})
	if err != nil {
		fmt.Printf("Error with migration Post. Error: %v", err.Error())
	}

	err = DB.AutoMigrate(&config.Statistic{})
	if err != nil {
		fmt.Printf("Error with migration Statistic. Error: %v", err.Error())
	}

	err = DB.AutoMigrate(&config.StatisticFromExcel{})
	if err != nil {
		fmt.Printf("Error with migration StatisticFromExcel. Error: %v", err.Error())
	}

	err = DB.AutoMigrate(&config.Patient{})
	if err != nil {
		fmt.Printf("Error with migration Patient. Error: %v", err.Error())
	}

	err = DB.AutoMigrate(&config.Diagnose{})
	if err != nil {
		fmt.Printf("Error with migration Diagnose. Error: %v", err.Error())
	}

	err = DB.AutoMigrate(&config.Exam{})
	if err != nil {
		fmt.Printf("Error with migration Exam. Error: %v", err.Error())
	}

}
