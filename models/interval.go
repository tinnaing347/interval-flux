package models

import (
	"github.com/mitchellh/mapstructure"
)

type Interval struct {
	Demand             float32 `mapstructure:"demand"`
	Energy             float32 `mapsturcture:"energy"`
	Time               string  `mapsturcture:"time"`
	Unique_meter_seqid string  `mapsturcture:"unique_meter_seqid"`
}

func Serializer(columns []string, values []interface{}) map[string]interface{} {
	map_ := make(map[string]interface{})

	for i := 0; i < len(columns); i++ {
		map_[columns[i]] = values[i]
	}

	return map_
}

func NewInterval(columns []string, values []interface{}) *Interval {

	interval_map := Serializer(columns, values)

	var interval Interval

	err := mapstructure.Decode(interval_map, &interval)
	if err != nil {
		panic(err)
	}
	return &interval
}
