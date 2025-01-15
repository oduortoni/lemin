package utils

// SliceContainsString checks if a given string (s) is present in the provided slice of strings (arr).
func SliceContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
