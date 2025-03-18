package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries/storage"
	"knp_server/internal/delivery/http/handlers/medical"
	"knp_server/internal/models"
	"knp_server/internal/utils"
	"net/http"
	"strconv"
)

func CreateEquipmentHandler(c *gin.Context) {
	var equipment models.Equipment

	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := storage.CreateEquipment(equipment)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating equipment. Error: %v", err))
	} else {
		c.JSON(http.StatusOK, equipment)
	}

}

func GetEquipmentsHandler(c *gin.Context) {

	equipments, err := storage.GetEquipments()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, equipments)
	}

}

func GetEquipmentBySerNumberHandler(c *gin.Context) {

	equipments, err := storage.GetEquipmentBySerNumber(handlers.GetUrlParam(c, "serNumber"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, equipments)
	}

}

func UpdateEquipmentHandler(c *gin.Context) {
	var equipment models.Equipment

	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := storage.UpdateEquipment(equipment)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating equipment. Error: %v", err))
	} else {
		c.JSON(http.StatusOK, equipment)
	}

}

func CreateComputerHandler(c *gin.Context) {
	var req struct {
		EquipmentID     uint   `json:"equipment_id" binding:"required"`
		ProcessorID     uint   `json:"processor_id" binding:"required"`
		RAMModuleIDs    []uint `json:"ram_module_ids"`
		StorageDriveIDs []uint `json:"storage_drive_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	computer := models.Computer{
		EquipmentID: req.EquipmentID,
		ProcessorID: req.ProcessorID,
	}

	for _, ramID := range req.RAMModuleIDs {
		computer.RAMModules = append(computer.RAMModules, models.RAMModule{ID: ramID})
	}

	for _, storageID := range req.StorageDriveIDs {
		computer.StorageDrives = append(computer.StorageDrives, models.StorageDevice{ID: storageID})
	}

	if err := storage.CreateComputer(&computer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating computer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Computer created successfully", "computer": computer})
}

func GetComputerHandler(c *gin.Context) {
	computerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid computer ID"})
		return
	}

	computer, err := storage.GetComputerByID(uint(computerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Computer not found"})
		return
	}

	c.JSON(http.StatusOK, computer)
}

//
//import (
//	"github.com/gin-gonic/gin"
//	"knp_server/internal/database/postgresql"
//	"knp_server/internal/models"
//	"net/http"
//)
//
//func CreateNomenclature(c *gin.Context) {
//
//	var nomenclature models.Nomenclature
//
//	err := c.ShouldBindJSON(&nomenclature)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateNomenclature(nomenclature)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func GetNomenclatures(c *gin.Context) {
//
//	nomenclatures, err := postgresql.GetNomenclatures()
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nomenclatures)
//	}
//
//}
//
//func UpdateNomenclature(c *gin.Context) {
//
//	var monitors []models.Monitor
//
//	err := c.ShouldBindJSON(&monitors)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.UpdateMonitor(monitors)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func DeleteNomenclature(c *gin.Context) {
//
//	var monitor models.Monitor
//
//	err := c.ShouldBindJSON(&monitor)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateMonitor(monitor)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateMonitor(c *gin.Context) {
//
//	var monitor models.Monitor
//
//	err := c.ShouldBindJSON(&monitor)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateMonitor(monitor)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func GetMonitors(c *gin.Context) {
//
//	monitors, err := postgresql.GetMonitors()
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, monitors)
//	}
//
//}
//
//func UpdateMonitors(c *gin.Context) {
//
//	var monitors []models.Monitor
//
//	err := c.ShouldBindJSON(&monitors)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.UpdateMonitor(monitors)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func DeleteMonitor(c *gin.Context) {
//
//	var monitor models.Monitor
//
//	err := c.ShouldBindJSON(&monitor)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateMonitor(monitor)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
////COMPUTERS
//
//func CreateComputer(c *gin.Context) {
//
//	var computer models.Computer
//
//	err := c.ShouldBindJSON(&computer)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateComputer(computer)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func GetComputers(c *gin.Context) {
//
//	computers, err := postgresql.GetComputers()
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, computers)
//	}
//
//}
//
//func UpdateComputers(c *gin.Context) {
//
//	var computers []models.Computer
//
//	err := c.BindJSON(&computers)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	}
//
//	err = postgresql.UpdateComputers(computers)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"IsUpdated": "Successfully",
//		})
//	}
//}
//
//func DeleteComputer(c *gin.Context) {
//
//	var computers []models.Computer
//
//	err := c.BindJSON(&computers)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	}
//
//	err = postgresql.UpdateComputers(computers)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"IsUpdated": "Successfully",
//		})
//	}
//}
//
//func CreatePeriphery(c *gin.Context) {
//
//	var periphery models.Periphery
//
//	err := c.ShouldBindJSON(&periphery)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreatePeriphery(periphery)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateContract(c *gin.Context) {
//
//	var contract models.Contract
//
//	err := c.ShouldBindJSON(&contract)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateContract(contract)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateProcessor(c *gin.Context) {
//
//	var processor models.Processor
//
//	err := c.ShouldBindJSON(&processor)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateProcessor(processor)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateStorageDevice(c *gin.Context) {
//
//	var device models.StorageDevice
//
//	err := c.ShouldBindJSON(&device)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateStorageDevice(device)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateRAM(c *gin.Context) {
//
//	var ram models.RAM
//
//	err := c.ShouldBindJSON(&ram)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateRAM(ram)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func CreateResponsePerson(c *gin.Context) {
//
//	var person models.RespPerson
//
//	err := c.ShouldBindJSON(&person)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateResponsePerson(person)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
////REPAIRS
//
//func CreateRepair(c *gin.Context) {
//
//	var repair models.Repair
//
//	err := c.ShouldBindJSON(&repair)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateRepair(repair)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func GetRepairs(c *gin.Context) {
//
//	repairs, err := postgresql.GetRepairs()
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, repairs)
//	}
//
//}
//
//func UpdateRepair(c *gin.Context) {
//
//	var repair models.Repair
//
//	err := c.ShouldBindJSON(&repair)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateRepair(repair)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
//
//func DeleteRepair(c *gin.Context) {
//
//	var repair models.Repair
//
//	err := c.ShouldBindJSON(&repair)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	err = postgresql.CreateRepair(repair)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"Error": err.Error(),
//		})
//	} else {
//		c.JSON(http.StatusOK, nil)
//	}
//
//}
