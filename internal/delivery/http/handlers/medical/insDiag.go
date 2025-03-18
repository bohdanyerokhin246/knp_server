package handlers

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"net/http"
)

func CreateInstrumentalDiagnostic(c *gin.Context) {

	var err error
	var instrumentalDiagnosticRequest []models.InstrumentalDiagnosticRequest

	err = c.ShouldBindJSON(&instrumentalDiagnosticRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proceduresID, specialistIDs := queries.CreateInstrumentalDiagnostic(instrumentalDiagnosticRequest)

	c.JSON(http.StatusOK, gin.H{
		"procedures": len(proceduresID),
		"specialist": len(specialistIDs),
	})
}

func GetInstrumentalDiagnosticsBySpecialistCode(c *gin.Context) {
	specCode := c.Param("code")

	instrumentalDiagnostic, err := queries.GetInstrumentalDiagnosticsBySpecialistCode(specCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, instrumentalDiagnostic)
	}

}

func GetInstrumentalDiagnosticHandler(c *gin.Context) {

	instrumentalDiagnostic, err := queries.GetInstrumentalDiagnostic()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, instrumentalDiagnostic)
	}

}
