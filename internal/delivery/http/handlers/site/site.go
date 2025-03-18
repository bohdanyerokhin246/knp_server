package site

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"net/http"
)

func CreateMenu(c *gin.Context) {

	menu := models.Menu{}

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data incorrect"})
		return
	}

	err := queries.CreateMenu(menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, 0)
}

//func GetMenu(c *gin.Context) {
//
//	menus, err := queries.GetMenu()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": "Problem getting menu",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, menus)
//}

func CreatePage(c *gin.Context) {

	page := models.Page{}

	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data incorrect"})
		return
	}

	err := queries.CreatePage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, 0)
}

func GetPage(c *gin.Context) {

	pages, err := queries.GetPages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Username or password is incorrect",
		})
		return
	}

	c.JSON(http.StatusOK, pages)
}
