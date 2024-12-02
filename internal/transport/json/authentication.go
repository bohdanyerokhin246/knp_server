package json

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
	"os"
	"time"
)

func LoginHandler(c *gin.Context) {
	jwtKey := os.Getenv("jwtKey")
	if jwtKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT Key not configured"})
		return
	}

	user := config.User{}

	// Привязка данных JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	userFromDB, err := postgresql.GetUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		return
	}

	// Установка времени истечения токена
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &config.Claims{
		Role: userFromDB.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token generation error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
