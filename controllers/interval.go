package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/tinnaing347/interval-flux/models"
)

func GetAll(c *gin.Context) { // Get model if exist
	var intervals []models.Interval

	q := client.NewQuery("SELECT * FROM hourly_data", "ivdb", "")

	response, err := models.DB.Query(q)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// need to panic for error
	if response.Error() == nil {
		columns := response.Results[0].Series[0].Columns
		for i := 0; i < len(response.Results[0].Series); i++ {
			interval := models.NewInterval(columns, response.Results[0].Series[0].Values[i])
			intervals = append(intervals, *interval)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": intervals})
}
