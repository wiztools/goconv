package goconv

import "strconv"

func Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func Float32(s string) (float32, error) {
	t, err := strconv.ParseFloat(s, 32)
	return float32(t), err
}
