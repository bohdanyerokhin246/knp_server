package http

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/delivery/http/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handlers.LoginHandler)

		medical := v1.Group("/medical")
		{
			odkClasses := medical.Group("/odk")
			{
				odkClasses.GET("/classes", handlers.GetODKsHandler)
				odkClasses.GET("/classes/:id", handlers.GetODKByIDHandler)
			}

			diagnosesByODK := medical.Group("/diagnoses")
			{
				diagnosesByODK.GET("/odks", handlers.GetODKDiagnosesHandler)
				diagnosesByODK.GET("/odks/:id", handlers.GetDiagnoseByODKIdHandler)
			}

			specialists := medical.Group("/specialists")
			{
				specialists.GET("", handlers.GetSpecialistsHandler)
				specialists.GET("/:code", handlers.GetSpecialistByCodeHandler)
			}

			labTests := medical.Group("/labTests")
			{
				labTests.GET("", handlers.GetLabTestsHandler)
				labTests.GET("/:code", handlers.GetLabTestByCodeHandler)
			}

			consultation := medical.Group("/consultation")
			{
				consultation.GET("", handlers.GetConsultationsHandler)
				consultation.GET("/:code", handlers.GetConsultationsByCodeHandler)
				consultation.GET("/spec/:code", handlers.GetConsultationsBySpecialistCode)
				consultation.POST("/create", handlers.CreateConsultation)

			}

			procedures := medical.Group("/procedures")
			{
				procedures.GET("", handlers.GetProceduresHandler)
				procedures.GET("/:code", handlers.GetProceduresByCodeHandler)
				procedures.GET("/spec/:code", handlers.GetProceduresBySpecialistCode)
				procedures.POST("/create", handlers.CreateProcedure)

			}

			instrumentalDiagnostic := medical.Group("/instrumentalDiagnostic")
			{
				instrumentalDiagnostic.GET("", handlers.GetInstrumentalDiagnosticHandler)
				//instrumentalDiagnostic.GET("/:code", handlers.GetProceduresByCodeHandler)
				instrumentalDiagnostic.GET("/spec/:code", handlers.GetInstrumentalDiagnosticsBySpecialistCode)
				instrumentalDiagnostic.POST("/create", handlers.CreateInstrumentalDiagnostic)

			}
		}

		storage := v1.Group("/storage")
		{
			repair := storage.Group("/equipments")
			{
				repair.GET("", handlers.GetEquipmentsHandler)
				repair.GET("/:serNumber", handlers.GetEquipmentBySerNumberHandler)
				//repair.GET("/spec/:code", handlers.GetInstrumentalDiagnosticsBySpecialistCode)
				repair.POST("/create", handlers.CreateInstrumentalDiagnostic)

			}

			computer := storage.Group("/computers")
			{
				computer.POST("", handlers.CreateComputerHandler)
				computer.GET("/:id", handlers.GetComputerHandler)
			}
		}

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
