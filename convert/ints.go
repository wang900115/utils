package convert

import "strconv"

func FromInt64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func FromUintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}
