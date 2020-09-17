package interval

import (
	"bytes"

	"github.com/tinnaing347/interval-flux/query"
)

/// how much will have to re write if we change to flux

type IntervalFilterInput struct {
	UniqueMeterSeqid string `form:"unique_meter_seqid"` // XXX DONT EVER USE SPACE OR UNDERSCORE IN FIELD NAME
	StartDate        string `form:"start_date"`
	EndDate          string `form:"end_date"`
	ExactTime        string `form:"exact_time"`
}

// a very tedious way to make a query string
func (r *IntervalFilterInput) MakeQueryString(command string, measurement string) string {
	var b bytes.Buffer
	b.WriteString(command)
	b.WriteString(" FROM ")
	b.WriteString(measurement)

	check_empty := r == &IntervalFilterInput{}
	if check_empty {
		return b.String()
	}

	flag := true
	b, flag = query.HandleTagFieldQuery(b, "unique_meter_seqid", r.UniqueMeterSeqid, flag)
	b, flag = query.HandleDateQuery(b, r.ExactTime, "=", flag)
	b, flag = query.HandleDateQuery(b, r.StartDate, ">=", flag)
	b, _ = query.HandleDateQuery(b, r.EndDate, "<", flag)

	return b.String()
}
