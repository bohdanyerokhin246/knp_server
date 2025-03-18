package news

import "github.com/gin-gonic/gin"

func RegisterNewsRoutes(v1 *gin.RouterGroup) {
	news := v1.Group("/news")
	{
		news.POST("", CreateNewsHandler)
		//news.GET("/:id", GetNewsByIDHandler)
		news.GET("", GetNewsListHandler)
		news.DELETE("", DeleteNewsHandler)
	}
}
