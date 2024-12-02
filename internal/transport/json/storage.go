package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreateMonitor(c *gin.Context) {

	var monitor config.Monitor

	err := c.ShouldBindJSON(&monitor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateMonitor(monitor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func GetMonitors(c *gin.Context) {

	monitors, err := postgresql.GetMonitors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, monitors)
	}

}

func UpdateMonitors(c *gin.Context) {

	var monitor config.Monitor

	err := c.ShouldBindJSON(&monitor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateMonitor(monitor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func DeleteMonitor(c *gin.Context) {

	var monitor config.Monitor

	err := c.ShouldBindJSON(&monitor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateMonitor(monitor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

//COMPUTERS

func CreateComputer(c *gin.Context) {

	var computer config.Computer

	err := c.ShouldBindJSON(&computer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateComputer(computer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func GetComputers(c *gin.Context) {

	computers, err := postgresql.GetComputers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, computers)
	}

}

func UpdateComputers(c *gin.Context) {

	var computers []config.Computer

	err := c.BindJSON(&computers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}

	err = postgresql.UpdateComputers(computers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"IsUpdated": "Successfully",
		})
	}
}

func DeleteComputer(c *gin.Context) {

	var computers []config.Computer

	err := c.BindJSON(&computers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}

	err = postgresql.UpdateComputers(computers)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"IsUpdated": "Successfully",
		})
	}
}

func CreatePeriphery(c *gin.Context) {

	var periphery config.Periphery

	err := c.ShouldBindJSON(&periphery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreatePeriphery(periphery)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func CreateContract(c *gin.Context) {

	var contract config.Contract

	err := c.ShouldBindJSON(&contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateContract(contract)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func CreateProcessor(c *gin.Context) {

	var processor config.Processor

	err := c.ShouldBindJSON(&processor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateProcessor(processor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func CreateStorageDevice(c *gin.Context) {

	var device config.StorageDevice

	err := c.ShouldBindJSON(&device)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateStorageDevice(device)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func CreateRAM(c *gin.Context) {

	var ram config.RAM

	err := c.ShouldBindJSON(&ram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateRAM(ram)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func CreateResponsePerson(c *gin.Context) {

	var person config.RespPerson

	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateResponsePerson(person)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

//REPAIRS

func CreateRepair(c *gin.Context) {

	var repair config.Repair

	err := c.ShouldBindJSON(&repair)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateRepair(repair)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func GetRepairs(c *gin.Context) {

	var repair config.Repair

	err := c.ShouldBindJSON(&repair)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateRepair(repair)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func UpdateRepair(c *gin.Context) {

	var repair config.Repair

	err := c.ShouldBindJSON(&repair)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateRepair(repair)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func DeleteRepair(c *gin.Context) {

	var repair config.Repair

	err := c.ShouldBindJSON(&repair)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateRepair(repair)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}
