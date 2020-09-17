package interval

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/tinnaing347/interval-flux/models"
)

func GetIntervals(c *gin.Context) { // Get model if exist
	var req IntervalFilterInput

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	fmt.Println(req.MakeQueryString("SELECT *", "hourly_data"))
	q := client.NewQuery(req.MakeQueryString("SELECT *", "hourly_data"), "ivdb", "")

	response, err := models.DB.Query(q)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if response.Error() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Error()})
		return
	}

	if len(response.Results[0].Series) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []string{}})
		return
	}

	intervals := ParseInfluxDBResult(response.Results[0])

	c.JSON(http.StatusOK, gin.H{"data": intervals})
}

func ParseInfluxDBResult(result client.Result) []Interval {
	var intervals []Interval
	columns := result.Series[0].Columns //result.MarshalJSON does not work for some reason
	for i := 0; i < len(result.Series[0].Values); i++ {
		interval := NewInterval(columns, result.Series[0].Values[i])
		intervals = append(intervals, *interval)
	}
	return intervals
}

func CreateIntervals(c *gin.Context) {
	var input Interval

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags, fields, time_ := input.TagField()

	models.CreateBatchPoint("ivdb", "hourly_data", tags, fields, time_)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteIntervals(c *gin.Context) {
	var req IntervalFilterInput

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	q := client.NewQuery(req.MakeQueryString("DROP SERIES", "hourly_data"), "ivdb", "")

	response, err := models.DB.Query(q)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if response.Error() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": response.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
