package postgresql

import (
	"knp_server/internal/config"
)

func CreatePost(post config.Post) (uint, error) {

	// INSERT INTO posts.posts (post_body, created_by, updated_at, deleted_at, created_at, create_date, is_actual) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	result := DB.Omit("ID").Create(&post)

	return post.ID, result.Error
}

func GetPosts() ([]config.Post, error) {

	var posts []config.Post

	// SELECT * FROM posts.posts WHERE is_actual = true ORDER BY id desc
	err := DB.Order("id desc").Find(&posts)

	if err.Error != nil {
		return nil, err.Error
	}

	return posts, err.Error
}

func UpdatePost(post config.Post) error {

	// UPDATE posts.posts SET `is_actual`= $1 WHERE `id` = $2
	err := DB.Model(&post).Where("id = ?", post.ID).Update("is_actual", post.IsActual)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
