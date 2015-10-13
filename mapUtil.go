package jsonUtil

// EqualStringKeyInterfaceMaps finds out if two maps with key typed in string are identical.
func EqualStringKeyInterfaceMaps(map1 map[string]interface{}, map2 map[string]interface{}) bool {
	if len(map1) == len(map2) {
		for key1, value1 := range map1 {
			if !EqualValues(value1, map2[key1]) {
				return false
			}
		}
		return true
	}
	return false
}
