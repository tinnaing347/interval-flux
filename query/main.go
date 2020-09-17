package query

import (
	"bytes"
	"strings"
)

func HandleTagFieldQuery(b bytes.Buffer, key_str string, val_str string, flag bool) (bytes.Buffer, bool) {
	if len(val_str) > 0 {
		b.WriteString(DetermineFirstCondition(flag))
		b = HandleMultiTagValues(b, key_str, val_str)
		flag = false
	}
	return b, flag
}

// returns partial query string for time related queries
func HandleDateQuery(b bytes.Buffer, date_str string, compare_str string, flag bool) (bytes.Buffer, bool) {
	if len(date_str) > 0 {
		b.WriteString(DetermineFirstCondition(flag))
		b.WriteString(" time")
		b.WriteString(compare_str)
		b.WriteString(EncloseWithSingleQuotes(date_str))
		flag = false
	}
	return b, flag
}

// Encloses a string in single quotes for influxdb; influxdbv1.8 only uses single quotes; 2.0 uses double quotes
func EncloseWithSingleQuotes(str string) string {
	var b bytes.Buffer
	b.WriteString("'")
	b.WriteString(str)
	b.WriteString("'")
	return b.String()
}

// return partial query string for multiple tag values; i.e add OR for multiple tag values
func HandleMultiTagValues(b bytes.Buffer, tag_key string, tag_val_str string) bytes.Buffer {
	str_ls := strings.Split(tag_val_str, ",")
	or_str := " "
	for i := 0; i < len(str_ls); i++ {
		b.WriteString(or_str)
		b.WriteString(tag_key)
		b.WriteString("=")
		b.WriteString(EncloseWithSingleQuotes(str_ls[i]))
		or_str = " OR "
	}
	return b
}

// true if it is Frist condition and return where; else it is second query conidition and returns and
func DetermineFirstCondition(flag bool) string {
	if flag {
		return " WHERE"
	} else {
		return " AND"
	}
}
