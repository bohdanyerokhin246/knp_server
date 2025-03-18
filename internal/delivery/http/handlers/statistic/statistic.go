package statistic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knp_server/internal/database/postgresql/queries"
	"knp_server/internal/models"
	"net/http"
)

func CreateStatisticPatient(c *gin.Context) {

	var err error
	var statisticPatients []models.StatisticPatient

	err = c.ShouldBindJSON(&statisticPatients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = queries.CreateStatisticPatient(statisticPatients)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func GetUnitOfDoctor(c *gin.Context) {

	doctors, err := queries.GetDepartmentByDoctor()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, doctors)
	}
}

func CreateEmz(c *gin.Context) {

	var err error
	var emzs []models.EMZ

	//Getting array fo emz in json
	err = c.ShouldBindJSON(&emzs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	fmt.Println("JSON is correct")
	err = queries.CreateEmzs(emzs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	}

	fmt.Println("EMZ created")
	err = queries.CorrectionEMZPaymentActual()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	}
	fmt.Println("EMZ price corrected")
	c.JSON(http.StatusOK, 0)
}

func GetStatisticByPackage(c *gin.Context) {

	//user, _ := c.Get("role")
	//if user == "admin" {
	statistics, err := queries.GetStatisticByPackage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}
	//} else {
	//	c.JSON(http.StatusUnauthorized, nil)
	//}
}

func GetStatisticByUnit(c *gin.Context) {

	statistics, err := queries.GetStatisticByUnit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}

	//user, _ := c.Get("role")
	//
	//if user == "admin" {
	//	statistics, err := postgresql.GetStatisticByUnit()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"Error": err.Error(),
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, statistics)
	//	}
	//} else {
	//	c.JSON(http.StatusUnauthorized, nil)
	//}
}

func GetDynamicAll(c *gin.Context) {

	statistics, err := queries.GetStatisticAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}
}

func GetDynamicByDoctor(c *gin.Context) {

	statistics := queries.GetIncludedSummarizeStatistic()
	c.JSON(http.StatusOK, statistics)

}

func GetDynamicByPackage(c *gin.Context) {

	user, _ := c.Get("role")
	if user == "admin" {
		statistics, err := queries.GetStatisticByPackage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, statistics)
		}
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func GetDynamicByUnit(c *gin.Context) {
	statistics, err := queries.GetStatisticByUnit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}

	//user, _ := c.Get("role")
	//
	//if user == "admin" {
	//	statistics, err := postgresql.GetStatisticByUnit()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"Error": err.Error(),
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, statistics)
	//	}
	//} else {
	//	c.JSON(http.StatusUnauthorized, nil)
	//}
}
