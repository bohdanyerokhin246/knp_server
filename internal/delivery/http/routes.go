package http

import (
	"github.com/gin-gonic/gin"
	handlers2 "knp_server/internal/delivery/http/handlers/medical"
	"knp_server/internal/delivery/http/handlers/news"
	"knp_server/internal/delivery/http/handlers/site"
	"knp_server/internal/delivery/http/handlers/statistic"
	"knp_server/internal/delivery/http/handlers/storage"
	"knp_server/internal/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.Use(middleware.JWTMiddleware())
	{
		menu := v1.Group("/menu")
		{
			menu.GET("", GetMenu)
		}

		v1.POST("/login", site.LoginHandler)

		medical := v1.Group("/medical")
		{
			odkClasses := medical.Group("/odk")
			{
				odkClasses.GET("/classes", handlers2.GetODKsHandler)
				odkClasses.GET("/classes/:id", handlers2.GetODKByIDHandler)
			}

			diagnosesByODK := medical.Group("/diagnoses")
			{
				diagnosesByODK.GET("/odks", handlers2.GetODKDiagnosesHandler)
				diagnosesByODK.GET("/odks/:id", handlers2.GetDiagnoseByODKIdHandler)
			}

			specialists := medical.Group("/specialists")
			{
				specialists.GET("", handlers2.GetSpecialistsHandler)
				specialists.GET("/:code", handlers2.GetSpecialistByCodeHandler)
			}

			labTests := medical.Group("/labTests")
			{
				labTests.GET("", handlers2.GetLabTestsHandler)
				labTests.GET("/:code", handlers2.GetLabTestByCodeHandler)
			}

			consultation := medical.Group("/consultations")
			{
				consultation.GET("", handlers2.GetConsultationsHandler)
				consultation.GET("/:code", handlers2.GetConsultationsByCodeHandler)
				consultation.GET("/spec/:code", handlers2.GetConsultationsBySpecialistCode)
				consultation.POST("/create", handlers2.CreateConsultation)

			}

			procedures := medical.Group("/procedures")
			{
				procedures.GET("", handlers2.GetProceduresHandler)
				procedures.GET("/:code", handlers2.GetProceduresByCodeHandler)
				procedures.GET("/spec/:code", handlers2.GetProceduresBySpecialistCode)
				procedures.POST("/create", handlers2.CreateProcedure)

			}

			instrumentalDiagnostic := medical.Group("/instrumentalDiagnostic")
			{
				instrumentalDiagnostic.GET("", handlers2.GetInstrumentalDiagnosticHandler)
				//instrumentalDiagnostic.GET("/:code", handlers.GetProceduresByCodeHandler)
				instrumentalDiagnostic.GET("/spec/:code", handlers2.GetInstrumentalDiagnosticsBySpecialistCode)
				instrumentalDiagnostic.POST("/create", handlers2.CreateInstrumentalDiagnostic)

			}
		}

		file := v1.Group("/upload")
		{
			file.POST("", news.UploadFileHandler)
		}

		statistic.RegisterStatisticRoutes(v1)
		storage.RegisterStorageRoutes(v1)
		news.RegisterNewsRoutes(v1)

	}

	//registerMenuRoutes(router)
	//registerPageRoutes(router)
	//registerPostRoutes(router)
	//registerStatisticRoutes(router)
	//registerFlgRoutes(router)
	//registerMonitorRoutes(router)
	//registerNomenclatureRoutes(router)
	//registerComputerRoutes(router)
	//registerRepairRoutes(router)
}

//
//func registerMenuRoutes(router *gin.Engine) {
//	menu := router.Group("/menu")
//	{
//		menu.POST("", handlers.CreateMenu)
//		menu.GET("/", handlers.GetMenu)
//	}
//}
//
//func registerPageRoutes(router *gin.Engine) {
//	page := router.Group("/page")
//	{
//		page.POST("/", handlers.CreatePage)
//		page.GET("/", handlers.GetPage)
//	}
//}
//
//func registerPostRoutes(router *gin.Engine) {
//	posts := router.Group("/posts")
//	{
//		posts.GET("/", handlers.GetPosts)
//		posts.POST("/", handlers.CreatePost)
//		posts.PUT("/", handlers.UpdatePost)
//	}
//}
//
//func registerStatisticRoutes(router *gin.Engine) {
//	statistic := router.Group("/statistic")
//	{
//		statistic.GET("/departments", handlers.GetUnitOfDoctor)
//		statistic.POST("/emzs", handlers.CreateEmz)
//		statistic.POST("/patients", handlers.CreateStatisticPatient)
//		statistic.GET("/byPackage", handlers.GetStatisticByPackage)
//		statistic.GET("/byUnit", handlers.GetStatisticByUnit)
//	}
//}
//
//func registerFlgRoutes(router *gin.Engine) {
//	flg := router.Group("/flg")
//	{
//		flg.POST("/patients", handlers.CreatePatient)
//		flg.POST("/diagnoses", handlers.CreateDiagnose)
//		flg.GET("/exams", handlers.GetExams)
//		flg.POST("/exams", handlers.CreateExam)
//		flg.POST("/therapists", handlers.CreateTherapist)
//	}
//}
//
//func registerMonitorRoutes(router *gin.Engine) {
//	monitors := router.Group("/monitors")
//	{
//		monitors.GET("/", handlers.GetMonitors)
//		monitors.POST("/", handlers.CreateMonitor)
//		monitors.PUT("/", handlers.UpdateMonitors)
//	}
//}
//
//func registerNomenclatureRoutes(router *gin.Engine) {
//	nomenclature := router.Group("/nomenclature")
//	{
//		nomenclature.GET("/", handlers.GetNomenclatures)
//		nomenclature.POST("/", handlers.CreateNomenclature)
//	}
//}
//
//func registerComputerRoutes(router *gin.Engine) {
//	computers := router.Group("/computers")
//	{
//		computers.POST("/", handlers.CreateComputer)
//		computers.GET("/", handlers.GetComputers)
//		computers.PUT("/", handlers.UpdateComputers)
//	}
//}
//
//func registerRepairRoutes(router *gin.Engine) {
//	repairs := router.Group("/repairs")
//	{
//		repairs.GET("/", handlers.GetRepairs)
//		repairs.POST("/", handlers.CreateRepair)
//	}
//}
