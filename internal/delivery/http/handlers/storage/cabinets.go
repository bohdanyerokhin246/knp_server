package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries/storage"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
)

func CreateCabinetHandler(c *gin.Context) {
	var equipment models.Equipment

	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not correct request"})
		return
	}
	err := storage.CreateEquipment(equipment)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating cabinet. Error: %v", err))
	} else {
		c.JSON(http.StatusOK, equipment)
	}
}

func GetCabinetsHandler(c *gin.Context) {
	cabinets, err := storage.GetCabinets()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, cabinets)
	}
}
