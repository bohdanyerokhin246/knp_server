package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
	"strconv"
)

//ODK classes

func GetODKsHandler(c *gin.Context) {

	var odks []models.ODK
	if err := queries.FetchData(&odks); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error getting ODK classes. Error: %v", err))
		return
	}
	c.JSON(http.StatusOK, odks)
}

func GetODKByIDHandler(c *gin.Context) {
	var odk models.ODK

	id, requestError := strconv.Atoi(GetUrlParam(c, "id"))
	if requestError != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Incorrect id. Error: %v", requestError))
		return
	}

	tx := queries.FetchDataByID(&odk, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("id = ?", uint(id)).
			Limit(1)
	})

	if tx.Error != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Error getting ODK classes with ID = %d. Error: %v", id, tx.Error))
		return
	}

	if tx.RowsAffected == 0 {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Record with ID = %d not found.", id))
		return
	}
	c.JSON(http.StatusOK, odk)
}

//Specialists

func GetSpecialistsHandler(c *gin.Context) {

	var medicalSpecialists []models.Specialist

	err := queries.FetchData(&medicalSpecialists, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("status = ?", true).
			Order("name asc")
	})

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Error getting Specialists. Error: %v", err),
		)
		return
	}
	c.JSON(http.StatusOK, medicalSpecialists)
}

func GetSpecialistByCodeHandler(c *gin.Context) {
	var specialist models.Specialist

	code := GetUrlParam(c, "code")

	tx := queries.FetchDataByID(&specialist, func(db *gorm.DB) *gorm.DB {
		return db.
			Where("code = ?", code).
			Limit(1)
	})

	if tx.Error != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Error getting Specialist with Code = %s. Error: %v", code, tx.Error))
		return
	}

	if tx.RowsAffected == 0 {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Record with Code = %s not found.", code))
		return
	}
	c.JSON(http.StatusOK, specialist)
}

//DiagnosesByODK

func GetODKDiagnosesHandler(c *gin.Context) {

	odks, err := queries.GetODKDiagnoses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error receiving data"})
		return
	}
	c.JSON(http.StatusOK, odks)
}

func GetDiagnoseByODKIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(GetUrlParam(c, "id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	odk, err := queries.GetODKDiagnoseById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ODK not found"})
		return
	}

	c.JSON(http.StatusOK, odk)
}

//Procedures

func GetProceduresHandler(c *gin.Context) {

	procedures, err := queries.GetProcedures()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func GetProceduresByCodeHandler(c *gin.Context) {

	procedures, err := queries.GetProceduresByCode(GetUrlParam(c, "code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func GetProceduresBySpecialistCode(c *gin.Context) {
	procedures, err := queries.GetProceduresBySpecialistCode(GetUrlParam(c, "code"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, procedures)
	}

}

func CreateProcedure(c *gin.Context) {

	var err error
	var proceduresRequest []models.Request

	err = c.ShouldBindJSON(&proceduresRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proceduresID, odkIDs, specialistIDs := queries.CreateProcedure(proceduresRequest)

	c.JSON(http.StatusOK, gin.H{
		"proceduresID":  proceduresID,
		"odkIDs":        odkIDs,
		"specialistIDs": specialistIDs,
	})
}

func GetUrlParam(c *gin.Context, key string) string {
	urlParam := c.Param(key)
	return urlParam
}
