package app

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql"
	"knp_server/internal/middleware"
	"knp_server/internal/transport/json"
)

func Run() {

	//Init router and postgresql
	r := gin.Default()
	postgresql.Connect()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.AuthMiddleware)

	r.POST("/login", json.Login)

	r.POST("/posts", json.CreatePost)
	r.GET("/posts", json.GetPosts)
	r.PUT("/posts", json.UpdatePost)

	r.GET("/statistics", json.GetAllEmz)
	r.GET("/statistics/included/summarized", json.GetIncludedSummarizeStatistic)
	r.GET("/statistic/departments", json.GetUnitOfDoctor)
	r.POST("/statistic/emzs", json.CreateEmz)
	r.POST("/statistic/patients", json.CreateStatisticPatient)
	r.GET("/statistics/summarize", json.GetSummarizeStatistic)

	r.POST("/flg/patients", json.CreatePatient)
	r.POST("/flg/diagnoses", json.CreateDiagnose)
	r.GET("/flg/exams", json.GetExams)
	r.POST("/flg/exams", json.CreateExam)
	r.POST("/flg/therapists", json.CreateTherapist)

	r.GET("/monitors", json.GetMonitors)
	r.POST("/monitors", json.CreateMonitor)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/contracts", json.CreateContract)
	//r.GET("/monitor/get", json.GetMonitors)
	//r.GET("/monitor/update", json.GetMonitors)

	r.POST("/periphery/create", json.CreatePeriphery)
	//r.GET("/periphery/get", json.GetMonitors)
	//r.GET("/periphery/update", json.GetMonitors)

	r.POST("/computer/create", json.CreateComputer)
	r.GET("/computer/get", json.GetComputers)
	//r.GET("/monitor/update", json.GetMonitors)

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

	_ = r.Run(":8081")
}
