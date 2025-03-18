package queries

import (
	"knp_server/internal/database/postgresql"
	"knp_server/internal/models"
)

func GetUser(login models.LoginRequest) (models.User, error) {

	var user models.User

	err := postgresql.DB.User.Where("username = ? AND password = ?", login.Username, login.Password).Find(&user)

	if err.Error != nil {
		return user, err.Error
	}

	return user, nil
}
