package utils

// SliceContainsSlice checks if any element from the second slice (arr2) is present in the first slice (arr1).
func SliceContainsSlice(arr1, arr2 []string) bool {
	for _, v := range arr2 {
		if SliceContainsString(arr1, v) { // would just compare direct to avoid the overhead of a function call
			return true
		}
	}
	return false
}
