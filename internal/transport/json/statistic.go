package json

import (
	"github.com/gin-gonic/gin"
	"knp_server/internal/config"
	"knp_server/internal/database/postgresql"
	"net/http"
)

func CreateStatisticPatient(c *gin.Context) {

	var err error
	var statisticPatients []config.StatisticPatient

	err = c.ShouldBindJSON(&statisticPatients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postgresql.CreateStatisticPatient(statisticPatients)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func GetUnitOfDoctor(c *gin.Context) {

	doctors, err := postgresql.GetDepartmentByDoctor()

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
	var emzs []config.EMZ

	//Getting array fo emz in json
	err = c.ShouldBindJSON(&emzs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	err = postgresql.CreateEmzs(emzs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	}

	err = postgresql.CorrectionEMZPaymentActual()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
	}
	c.JSON(http.StatusOK, 0)
}

func GetStatisticAll(c *gin.Context) {

	statistics, err := postgresql.GetStatisticAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}
}

func GetStatisticByDoctor(c *gin.Context) {
	//var err error
	user, _ := c.Get("role")
	if user == "admin" {
		statistics := postgresql.GetIncludedSummarizeStatistic()
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"Error": err.Error(),
		//	})
		//} else {
		c.JSON(http.StatusOK, statistics)
		//}
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func GetStatisticByPackage(c *gin.Context) {

	user, _ := c.Get("role")
	if user == "admin" {
		statistics, err := postgresql.GetStatisticByPackage()
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

func GetStatisticByUnit(c *gin.Context) {
	user, _ := c.Get("role")

	if user == "admin" {
		statistics, err := postgresql.GetStatisticByUnit()
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

func GetDynamicAll(c *gin.Context) {

	statistics, err := postgresql.GetStatisticAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, statistics)
	}
}

func GetDynamicByDoctor(c *gin.Context) {
	//var err error
	user, _ := c.Get("role")
	if user == "admin" {
		statistics := postgresql.GetIncludedSummarizeStatistic()
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"Error": err.Error(),
		//	})
		//} else {
		c.JSON(http.StatusOK, statistics)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func GetDynamicByPackage(c *gin.Context) {

	user, _ := c.Get("role")
	if user == "admin" {
		statistics, err := postgresql.GetStatisticByPackage()
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
	user, _ := c.Get("role")

	if user == "admin" {
		statistics, err := postgresql.GetStatisticByUnit()
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
