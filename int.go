package goconv

import "strconv"

func Int64(s string) (int64, error) {
	if result, err := strconv.ParseInt(s, 10, 64); err != nil {
		return 0, err
	} else {
		return result, nil
	}
}

func Uint64(s string) (uint64, error) {
	if result, err := strconv.ParseUint(s, 10, 64); err != nil {
		return 0, err
	} else {
		return result, nil
	}
}
