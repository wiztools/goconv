package goconv

func Uniq[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
