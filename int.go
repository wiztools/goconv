package goconv

import "strconv"

func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func Int32(s string) (int32, error) {
	t, err := strconv.ParseInt(s, 10, 32)
	return int32(t), err
}

func Uint32(s string) (uint32, error) {
	t, err := strconv.ParseUint(s, 10, 32)
	return uint32(t), err
}

func Int16(s string) (int16, error) {
	t, err := strconv.ParseInt(s, 10, 16)
	return int16(t), err
}

func Uint16(s string) (uint16, error) {
	t, err := strconv.ParseUint(s, 10, 16)
	return uint16(t), err
}

func Int8(s string) (int8, error) {
	t, err := strconv.ParseInt(s, 10, 8)
	return int8(t), err
}

func Uint8(s string) (uint8, error) {
	t, err := strconv.ParseUint(s, 10, 8)
	return uint8(t), err
}
