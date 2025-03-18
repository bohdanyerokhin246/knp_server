package handlers

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"net/http"
)

//ODK classes

func GetConsultationsHandler(c *gin.Context) {

	procedures, err := queries.GetConsultations()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func GetConsultationsByCodeHandler(c *gin.Context) {

	procedures, err := queries.GetProceduresByCode(c.Param("code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func GetConsultationsBySpecialistCode(c *gin.Context) {
	procedures, err := queries.GetProceduresBySpecialistCode(c.Param("code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func CreateConsultation(c *gin.Context) {

	var err error
	var consultationRequest []models.Request

	err = c.ShouldBindJSON(&consultationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consultationID, consultationODKID, consultationSpecialistID := queries.CreateConsultation(consultationRequest)

	c.JSON(http.StatusOK, gin.H{
		"consultationID":           consultationID,
		"consultationODKID":        consultationODKID,
		"consultationSpecialistID": consultationSpecialistID,
	})
}
