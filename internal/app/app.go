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

	r.GET("/statistics/get", json.GetStatistics)

	_ = r.Run(":8081")
}
