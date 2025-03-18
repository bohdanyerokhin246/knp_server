package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

type DBConnections struct {
	Site      *gorm.DB
	User      *gorm.DB
	Post      *gorm.DB
	Statistic *gorm.DB
	Medical   *gorm.DB
	Storage   *gorm.DB
	MedInfo   *gorm.DB
}

var DB *DBConnections

func Connect() error {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("hostDB"),
		os.Getenv("portDB"),
		os.Getenv("userDB"),
		os.Getenv("passwordDB"),
		os.Getenv("nameDB"),
		os.Getenv("sslModeDB"))

	schemas := map[string]string{
		"Site":      "site.",
		"User":      "user.",
		"News":      "post.",
		"Statistic": "statistic.",
		"Medical":   "flg.",
		"Storage":   "storage.",
		"MedInfo":   "medical.",
	}

	dbConns := &DBConnections{}

	for key, schemaName := range schemas {
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{TablePrefix: schemaName},
		})
		if err != nil {
			return fmt.Errorf("failed to connect to %s: %w", key, err)
		}

		switch key {
		case "Site":
			dbConns.Site = db
		case "User":
			dbConns.User = db
		case "News":
			dbConns.Post = db
		case "Statistic":
			dbConns.Statistic = db
		case "Medical":
			dbConns.Medical = db
		case "Storage":
			dbConns.Storage = db
		case "MedInfo":
			dbConns.MedInfo = db
		}
	}

	fmt.Println("Connection to DB successfully ")

	DB = dbConns
	return nil
}
