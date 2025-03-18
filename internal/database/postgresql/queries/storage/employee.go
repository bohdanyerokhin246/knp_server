package storage

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreateEmployee(employee models.Employee) error {
	err := postgresql.DB.Storage.Create(&employee)
	if err != nil {
		return err.Error
	}
	return nil
}

func GetEmployees() ([]models.Employee, error) {
	var employee []models.Employee
	if err := postgresql.DB.Storage.Find(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}
