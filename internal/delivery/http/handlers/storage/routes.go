package storage

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/middleware"
)

func RegisterStorageRoutes(v1 *gin.RouterGroup) {
	storage := v1.Group("/storage")
	storage.Use(middleware.AuthMiddleware("admin"))
	{
		equipment := storage.Group("/equipments")
		{
			equipment.GET("", GetEquipmentsHandler)
			equipment.GET("/sn/:serNumber", GetEquipmentBySerNumberHandler)
			//repair.GET("/spec/:code", handlers.GetInstrumentalDiagnosticsBySpecialistCode)
			equipment.POST("", CreateEquipmentHandler)
			equipment.PUT("", UpdateEquipmentHandler)

		}

		computer := storage.Group("/computers")
		{
			computer.POST("", CreateComputerHandler)
			computer.GET("/:id", GetComputerHandler)
		}

		cabinet := storage.Group("/cabinets")
		{
			//cabinet.POST("", CreateComputerHandler)
			cabinet.GET("", GetCabinetsHandler)
		}

		employee := storage.Group("/employees")
		{
			//employee.POST("", CreateComputerHandler)
			employee.GET("", GetEmployeesHandler)
		}
	}
}
