package helper

import (
	"time"
	"webhook/variable"
)

func ParseLocalTime(from string) (res int64) {
	loc, _ := WIB()
	t, _ := time.ParseInLocation(variable.TimeLayout, from, loc)

	return t.UTC().Unix()
}
