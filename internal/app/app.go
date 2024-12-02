package app

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/transport/json"
	"net/http"
)

func Run() {

	//Init router and postgresql
	r := gin.Default()
	postgresql.Connect()

	// Пример защищённого маршрута
	protected := r.Group("/protected")

	protected.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.AuthMiddleware)

	r.GET("/test", json.GetTest)
	r.POST("/test", json.CreateTest)

	r.POST("/login", json.LoginHandler)

	r.POST("/menu", json.CreateMenu)
	r.GET("/menu", json.GetMenu)
	r.POST("/page", json.CreatePage)
	r.GET("/page", json.GetPage)

	r.POST("/posts", json.CreatePost)
	r.GET("/posts", json.GetPosts)
	r.PUT("/posts", json.UpdatePost)

	//Statistic

	r.GET("/statistic/departments", json.GetUnitOfDoctor)
	r.POST("/statistic/emzs", json.CreateEmz)
	r.POST("/statistic/patients", json.CreateStatisticPatient)

	r.GET("/statistics/full", json.GetStatisticAll)
	r.GET("/statistics/byDoctor", json.GetStatisticByDoctor)
	r.GET("/statistics/byPackage", json.GetStatisticByPackage)
	r.GET("/statistics/byUnit", json.GetStatisticByUnit)

	r.GET("/dynamics/full", json.GetDynamicAll)
	r.GET("/dynamics/byDoctor", json.GetDynamicByDoctor)
	r.GET("/dynamics/byPackage", json.GetDynamicByPackage)
	r.GET("/dynamics/byUnit", json.GetDynamicByUnit)

	/////////////

	r.POST("/flg/patients", json.CreatePatient)
	r.POST("/flg/diagnoses", json.CreateDiagnose)
	r.GET("/flg/exams", json.GetExams)
	r.POST("/flg/exams", json.CreateExam)
	r.POST("/flg/therapists", json.CreateTherapist)

	r.GET("/monitors", json.GetMonitors)
	r.POST("/monitors", json.CreateMonitor)
	r.PUT("/monitors", json.UpdateMonitors)

	r.POST("/contracts", json.CreateContract)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/periphery/create", json.CreatePeriphery)
	//r.GET("/periphery/get", json.GetMonitors)
	//r.GET("/periphery/update", json.GetMonitors)

	r.POST("/computer", json.CreateComputer)
	r.GET("/computers", json.GetComputers)
	r.PUT("/computers", json.UpdateComputers)

	r.POST("/processor/create", json.CreateProcessor)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/storageDevice/create", json.CreateStorageDevice)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/ram/create", json.CreateRAM)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/responsePerson/create", json.CreateResponsePerson)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/repairs", json.CreateRepair)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	_ = r.Run(":8081")
}
