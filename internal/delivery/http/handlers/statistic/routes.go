package statistic

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/middleware"
)

func RegisterStatisticRoutes(v1 *gin.RouterGroup) {
	stat := v1.Group("/statistics")
	stat.Use(middleware.AuthMiddleware("admin", "direct"))
	{
		stat.GET("/departments", GetUnitOfDoctor)
		stat.POST("/emzs", CreateEmz)
		stat.POST("/patients", CreateStatisticPatient)
		//
		stat.GET("/byDoctor", GetDynamicByDoctor)
		//
		stat.GET("/byPackage", GetStatisticByPackage).Use(middleware.AuthMiddleware("admin", "direct"))
		stat.GET("/byUnit", GetStatisticByUnit).Use(middleware.AuthMiddleware("admin", "direct"))
	}
}
