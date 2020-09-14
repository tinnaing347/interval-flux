package models

import (
	"github.com/mitchellh/mapstructure"
)

type Interval struct {
	Demand             float32 `mapstructure:"demand" json:"demand"`
	Energy             float32 `mapsturcture:"energy" json:"energy"`
	Time               string  `mapsturcture:"time" json:"time"`
	Unique_meter_seqid string  `mapsturcture:"unique_meter_seqid" json:"unique_meter_seqid"`
}

func (i *Interval) TagField() (map[string]string, map[string]interface{}) {

	tags := map[string]string{"unique_meter_seqid": i.Unique_meter_seqid}
	fields := map[string]interface{}{
		"energy": i.Energy,
		"demand": i.Demand,
		"time":   i.Time,
	}

	return tags, fields
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

	if err := mapstructure.Decode(interval_map, &interval); err != nil {
		panic(err)
	}
	return &interval
}
