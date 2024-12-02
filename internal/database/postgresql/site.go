package postgresql

import (
	"knp_server/internal/config"
)

func CreateMenu(menu config.Menu) error {

	result := DBSite.Create(&menu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMenu() ([]config.Menu, error) {

	var menus []config.Menu

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DBSite.Find(&menus)

	if err.Error != nil {
		return nil, err.Error
	}

	return menus, err.Error
}

func CreatePage(page config.Page) error {

	result := DBSite.Create(&page)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPages() ([]config.Page, error) {

	var pages []config.Page

	err := DBSite.Find(&pages)
	if err.Error != nil {
		return nil, err.Error
	}

	return pages, err.Error
}
