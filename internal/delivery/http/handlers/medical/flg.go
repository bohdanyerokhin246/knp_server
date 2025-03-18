package handlers

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"net/http"
)

func CreatePatient(c *gin.Context) {

	var patients []models.Patient
	err := c.ShouldBindJSON(&patients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queries.CreatePatient(patients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateDiagnose(c *gin.Context) {

	var diagnoses []models.DiagnoseFLG
	err := c.ShouldBindJSON(&diagnoses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queries.CreateDiagnose(diagnoses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateTherapist(c *gin.Context) {

	var therapists []models.Therapist
	err := c.ShouldBindJSON(&therapists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queries.CreateTherapist(therapists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateExam(c *gin.Context) {

	var exams []models.Exam
	err := c.ShouldBindJSON(&exams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queries.CreateExam(exams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetExams(c *gin.Context) {
	exams, err := queries.GetExams()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, exams)
	}

}
