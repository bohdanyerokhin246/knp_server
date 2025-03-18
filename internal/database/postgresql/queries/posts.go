package queries

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func CreatePost(post models.News) (uint, error) {

	result := postgresql.DB.Post.Create(&post)

	return post.ID, result.Error
}

func GetNews() ([]models.News, error) {

	var posts []models.News

	err := postgresql.DB.Post.Order("id desc").Find(&posts)

	if err.Error != nil {
		return nil, err.Error
	}

	return posts, err.Error
}

func UpdatePost(post models.News) error {

	err := postgresql.DB.Post.Model(&post).Where("id = ?", post.ID).Update("is_actual", post.IsActual)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeleteNews(post models.News) error {

	err := postgresql.DB.Post.Where("id = ?", post.ID).Delete(&post)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
