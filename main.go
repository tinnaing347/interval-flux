package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tinnaing347/interval-flux/interval"
	"github.com/tinnaing347/interval-flux/models"
)

//_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
//client "github.com/influxdata/influxdb1-client/v2"

func main() {

	router := gin.New()
	r := router.Group("/v1")

	models.CreateClient(os.Getenv("DB_ADDR"))

	defer models.DB.Close()

	r.GET("/intervals", interval.GetIntervals)
	r.POST("/intervals", interval.CreateIntervals)
	r.DELETE("/intervals", interval.DeleteIntervals)
	router.Run()

	//client code for influxdb v2.0
	// userName := "admin"
	// password := "admin"
	// client := influxdb2.NewClient("http://localhost:8086", fmt.Sprintf("%s:%s", userName, password))

	// queryAPI := client.QueryAPI("")

	// result, err := queryAPI.Query(context.Background(), `from(bucket:"ivdb")|> range(start: 2020-08-22T23:30:00Z) |> filter(fn: (r) => r._measurement == "hourly_data")`)

	// // fmt.Println(result.Results)
	// if err == nil {
	// 	for result.Next() {
	// 		fmt.Println(result.Record())
	// 		fmt.Println(result.Record())
	// 	}
	// 	if result.Err() != nil {
	// 		fmt.Printf("Query error: %s\n", result.Err().Error())
	// 	}
	// } else {
	// 	fmt.Printf("Query error: %s\n", err.Error())
	// }
	// // Close client
	// client.Close()
}
