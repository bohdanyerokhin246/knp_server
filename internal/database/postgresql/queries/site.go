package queries

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreateMenu(menu models.Menu) error {

	result := postgresql.DB.Site.Create(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMenu(role string) ([]models.MenuItem, error) {

	var menu []models.MenuItem

	err := postgresql.DB.Site.Preload("Submenu").Where("parent_id IS NULL AND roles LIKE ?", "%"+role+"%").Find(&menu).Error
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func CreatePage(page models.Page) error {

	result := postgresql.DB.Site.Create(&page)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPages() ([]models.Page, error) {

	var pages []models.Page

	err := postgresql.DB.Site.Find(&pages)
	if err.Error != nil {
		return nil, err.Error
	}

	return pages, err.Error
}
