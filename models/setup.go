package models

import (
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

var DB client.Client

func CreateClient(db_addr string) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: db_addr,
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = c
}

func CreateBatchPoint(db string, measurement string, tags map[string]string, fields map[string]interface{}, time_ time.Time) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	pt, err := client.NewPoint(measurement, tags, fields, time_)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := DB.Write(bp); err != nil {
		log.Fatal(err)
	}
}
