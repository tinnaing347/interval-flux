package interval

import (
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Interval struct {
	UniqueMeterSeqid string  `mapstructure:"unique_meter_seqid" json:"unique_meter_seqid"`
	Demand           float32 `mapstructure:"demand" json:"demand"`
	Energy           float32 `mapsturcture:"energy" json:"energy"`
	Time             string  `mapsturcture:"time" json:"time"`
}

func (i *Interval) TagField() (map[string]string, map[string]interface{}, time.Time) {

	tags := map[string]string{"unique_meter_seqid": i.UniqueMeterSeqid}
	fields := map[string]interface{}{
		"energy": i.Energy,
		"demand": i.Demand,
	}
	time_, err := time.Parse(time.RFC3339, i.Time) //2017-07-01T00:00:00Z, UTC if no timezone
	if err != nil {
		log.Fatal(err)
	}
	return tags, fields, time_
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
