package postgresql

import (
	"knp_server/internal/config"
)

func GetUser(user config.User) (config.User, error) {

	err := DBUser.Where("username = ?", user.Username).Where("password = ?", user.Password).Find(&user)

	if err.Error != nil {
		return user, err.Error
	}

	return user, nil
}
