package jsonUtil

// EqualStringKeyInterfaceMapSlices finds out if two slices are identical.
func EqualStringKeyInterfaceMapSlices(slice1 []map[string]interface{}, slice2 []map[string]interface{}) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if !EqualValues(x, slice2[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// EqualIntSlices finds out if two slices of strings are identical.
func EqualIntSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if x != slice2[i] {
				return false
			}
		}
		return true
	}
	return false
}

// EqualStringSlices finds out if two slices of strings are identical.
func EqualStringSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if x != slice2[i] {
				return false
			}
		}
		return true
	}
	return false
}
