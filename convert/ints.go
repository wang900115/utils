package convert

import "strconv"

func FromInt64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}
