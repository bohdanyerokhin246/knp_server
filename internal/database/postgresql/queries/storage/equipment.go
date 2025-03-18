package storage

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreateEquipment(equipment models.Equipment) error {
	err := postgresql.DB.Storage.Create(&equipment)
	if err != nil {
		return err.Error
	}
	return nil
}

func GetEquipments() ([]models.Equipment, error) {
	var equipments []models.Equipment
	err := postgresql.DB.Storage.
		Preload("Cabinet").
		Preload("Employee").
		Order("id asc").
		Find(&equipments).Error
	if err != nil {
		return nil, err
	}

	return equipments, nil
}

func GetEquipmentBySerNumber(code string) (models.Equipment, error) {
	var equipment models.Equipment
	if err := postgresql.DB.Storage.Where("inventory_number = ? OR serial_number = ?", code, code).Find(&equipment).Error; err != nil {
		return models.Equipment{}, err
	}
	return equipment, nil
}

func UpdateEquipment(equipment models.Equipment) error {
	err := postgresql.DB.Storage.Save(&equipment).Error
	if err != nil {
		return err
	}
	return nil
}
