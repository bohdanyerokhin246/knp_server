package postgresql

import "knp_server/internal/config"

func GetStatisticsOrderByDoctor() ([]config.Statistic, error) {

	var statistics []config.Statistic

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DB.Order("full_name asc").Find(&statistics)

	if err.Error != nil {
		return nil, err.Error
	}

	return statistics, err.Error
}

func GetStatisticsOrderByUnit() ([]config.Statistic, error) {

	var statistics []config.Statistic

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DB.Order("unit asc").Find(&statistics)

	if err.Error != nil {
		return nil, err.Error
	}

	return statistics, err.Error
}

func GetStatisticsOrderByPackage() ([]config.Statistic, error) {

	var statistics []config.Statistic

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DB.Order("package asc").Find(&statistics)

	if err.Error != nil {
		return nil, err.Error
	}

	return statistics, err.Error
}
