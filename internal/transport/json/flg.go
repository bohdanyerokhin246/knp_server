package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreatePatient(c *gin.Context) {

	var patients []config.Patient
	err := c.ShouldBindJSON(&patients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreatePatient(patients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateDiagnose(c *gin.Context) {

	var diagnoses []config.Diagnose
	err := c.ShouldBindJSON(&diagnoses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateDiagnose(diagnoses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateTherapist(c *gin.Context) {

	var therapists []config.Therapist
	err := c.ShouldBindJSON(&therapists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateTherapist(therapists)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func CreateExam(c *gin.Context) {

	var exams []config.Exam
	err := c.ShouldBindJSON(&exams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateExam(exams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetExams(c *gin.Context) {
	exams, err := postgresql.GetExams()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, exams)
	}

}
