package storage

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreateCabinet(equipment models.Equipment) (uint, error) {
	err := postgresql.DB.Storage.Create(&equipment)
	if err != nil {
		return 0, err.Error
	}
	return equipment.ID, nil
}

func GetCabinets() ([]models.Cabinet, error) {
	var cabinets []models.Cabinet
	if err := postgresql.DB.Storage.Find(&cabinets).Error; err != nil {
		return nil, err
	}
	return cabinets, nil
}

//func GetEquipmentBySerNumber(code string) (models.Equipment, error) {
//	var equipment models.Equipment
//
//	if err := postgresql.DB.Storage.Where("serial_number = ?", code).Find(&equipment).Error; err != nil {
//		return models.Equipment{}, err
//	}
//	return equipment, nil
//}
