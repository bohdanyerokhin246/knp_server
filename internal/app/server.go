package app

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/delivery/http"
	"knp_server/internal/middleware"
	"log"
)

func serverInitAndStart() {
	// Init Gin router
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	http.RegisterRoutes(router)

	//Starting server
	err := router.Run(":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
