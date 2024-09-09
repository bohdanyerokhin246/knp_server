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
	err := DB.Where("is_actual = ?", "true").Order("id desc").Find(&posts)

	if err.Error != nil {
		return nil, err.Error
	}

	return posts, err.Error
}
