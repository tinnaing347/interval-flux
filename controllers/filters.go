package controllers

import "bytes"

type IntervalFilterInput struct {
	UniqueMeterSeqid string `form:"unique_meter_seqid"` // XXX DONT EVER USE SPACE OR UNDERSCORE IN FIELD NAME
	StartDate        string `form:"start_date"`
	EndDate          string `form:"end_date"`
}

func (r *IntervalFilterInput) unique_meter_seqid(b bytes.Buffer) (bytes.Buffer, bool) {
	flag := false
	if len(r.UniqueMeterSeqid) > 0 {
		b.WriteString(" unique_meter_seqid = ")
		b.WriteString("'")
		b.WriteString(r.UniqueMeterSeqid)
		b.WriteString("'")
		flag = true
	}
	return b, flag
}

func (r *IntervalFilterInput) start_date(b bytes.Buffer, flag bool) (bytes.Buffer, bool) {
	if len(r.StartDate) > 0 {
		if flag {
			b.WriteString(" AND")
		}
		b.WriteString(" time <= ")
		b.WriteString(r.StartDate)
		flag = true
	}
	return b, flag
}

func (r *IntervalFilterInput) end_date(b bytes.Buffer, flag bool) bytes.Buffer {
	if len(r.EndDate) > 0 {
		if flag {
			b.WriteString(" AND")
		}
		b.WriteString(" time > ")
		b.WriteString(r.EndDate)
	}
	return b
}

// a very tedious way to make a query string
func (r *IntervalFilterInput) MakeQueryString(measurement string) string {
	var b bytes.Buffer
	b.WriteString("SELECT * FROM ")
	b.WriteString(measurement)

	check_empty := r == &IntervalFilterInput{}
	if check_empty {
		return b.String()
	}

	b, flag := r.unique_meter_seqid(b)
	b, flag = r.start_date(b, flag)
	b = r.end_date(b, flag)

	return b.String()
}
