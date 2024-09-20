package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreatePatient(c *gin.Context) {

	var patient config.Patient
	err := c.ShouldBindJSON(&patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CheckPatientUnique(patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
