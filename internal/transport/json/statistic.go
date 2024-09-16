package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func GetStatisticsOrderByDoctor(c *gin.Context) {

	statistics, err := postgresql.GetStatisticsOrderByDoctor()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Statistics": statistics,
		})
	}

}
func GetStatisticsOrderByUnit(c *gin.Context) {

	statistics, err := postgresql.GetStatisticsOrderByUnit()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Statistics": statistics,
		})
	}

}
func GetStatisticsOrderByPackage(c *gin.Context) {

	statistics, err := postgresql.GetStatisticsOrderByPackage()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Statistics": statistics,
		})
	}

}
