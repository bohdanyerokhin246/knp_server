package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func GetStatistics(c *gin.Context) {

	statistics, err := postgresql.GetStatistics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}

}
