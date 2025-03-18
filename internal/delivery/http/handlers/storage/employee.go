package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries/storage"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
)

func CreateEmployeeHandler(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request incorrect"})
		return
	}
	err := storage.CreateEmployee(employee)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating employee. Error: %v", err))
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func GetEmployeesHandler(c *gin.Context) {
	employees, err := storage.GetEmployees()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, employees)
	}
}
