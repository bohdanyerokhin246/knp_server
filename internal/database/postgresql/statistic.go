package postgresql

import "knp_server/internal/config"

func GetStatistics() ([]config.Statistic, error) {

	var statistics []config.Statistic

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DB.Find(&statistics)

	if err.Error != nil {
		return nil, err.Error
	}

	return statistics, err.Error
}
