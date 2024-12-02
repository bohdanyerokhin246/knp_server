package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreateMenu(c *gin.Context) {

	menu := config.Menu{}

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data incorrect"})
		return
	}

	err := postgresql.CreateMenu(menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, 0)
}

func GetMenu(c *gin.Context) {

	menus, err := postgresql.GetMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Problem getting menu",
		})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func CreatePage(c *gin.Context) {

	page := config.Page{}

	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data incorrect"})
		return
	}

	err := postgresql.CreatePage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, 0)
}

func GetPage(c *gin.Context) {

	pages, err := postgresql.GetPages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, pages)
}
