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
