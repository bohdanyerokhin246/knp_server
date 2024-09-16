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

	//JSON endpoint
	r.POST("/post/create", json.CreatePost)
	r.GET("/posts/get", json.GetPosts)
	r.POST("/post/update", json.UpdatePost)

	//r.POST("/post/create", json.CreatePost)
	r.GET("/statisticsOrderByDoctor/get", json.GetStatisticsOrderByDoctor)
	r.GET("/statisticsOrderByUnit/get", json.GetStatisticsOrderByUnit)
	r.GET("/statisticsOrderByPackage/get", json.GetStatisticsOrderByPackage)
	//r.POST("/post/update", json.UpdatePost)

	_ = r.Run(":8081")
}
