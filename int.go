package goconv

import "strconv"

func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
