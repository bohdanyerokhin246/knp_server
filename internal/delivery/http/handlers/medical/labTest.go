package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
)

func GetLabTestsHandler(c *gin.Context) {
	var labServices []models.LabTest
	if err := queries.FetchData(&labServices); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error getting ODK classes. Error: %v", err))
		return
	}
	c.JSON(http.StatusOK, labServices)
}

func GetLabTestByCodeHandler(c *gin.Context) {
	labServices, err := queries.GetLabTestByCode(GetUrlParam(c, "code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, labServices)
	}

}
