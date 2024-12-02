package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"knp_server/internal/config"
	"net/http"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authToken, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("authToken")
	if authHeader == "" {
		// Назначаем роль "user" и продолжаем выполнение для пользователей без токена
		c.Set("role", "user")
		return
	}

	// Извлечение токена из формата "Bearer <token>"
	var tokenString string
	if _, err := fmt.Sscanf(authHeader, "%s", &tokenString); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Неверный формат токена"})
		return
	}

	claims := &config.Claims{}

	// Получение ключа для подписи JWT
	jwtKey := []byte(os.Getenv("jwtKey"))
	if len(jwtKey) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT ключ не настроен"})
		return
	}

	// Валидация и парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		// Проверяем на истечение срока действия
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized,
				//gin.H{"message": "Expired token"},
				nil,
			)
		} else {
			c.JSON(http.StatusUnauthorized,
				//gin.H{"message": "Incorrect token"},
				nil,
			)
		}
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Неверный токен"})
		return
	}

	// Установка роли пользователя из токена в контекст
	c.Set("role", claims.Role)
	return
}
