package convert

import (
	"strconv"
	"strings"
	"time"
)

func FromStringToTimeStamp(s string) int64 {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func FromStringToBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "1":
		return true

	case "0":
		return false
	default:
		return false
	}
}

func FromStringToTimeTime(s string) time.Time {
	val, _ := time.Parse(time.RFC3339, s)
	return val
}

func FromStringToint(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
