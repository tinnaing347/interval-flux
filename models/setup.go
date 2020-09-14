package models

import (
	"log"
	"os"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/joho/godotenv"
)

var DB client.Client

func CreateClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: os.Getenv("API_ADDRESS"), //make a setting file instead?/
	})

	DB = c
}

func CreateBatchPoint(db string, measurement string, tags map[string]string, fields map[string]interface{}) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	pt, err := client.NewPoint(measurement, tags, fields)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := DB.Write(bp); err != nil {
		log.Fatal(err)
	}
}
