package jsonUtil

import (
	"fmt"
	"reflect"
)

// EqualInterfaces compares two interface{}s' equal. Currently targeted at value in bool/string/float64/reflect.Kind and complex value map/slice derived from them.
func EqualInterfaces(value1 interface{}, value2 interface{}) bool {
	if value1 == nil || value2 == nil {
		return false
	}
	switch valueA := value1.(type) {
	case bool:
		valueB, ok := value2.(bool)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case float64:
		valueB, ok := value2.(float64)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case string:
		valueB, ok := value2.(string)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case map[string]interface{}:
		valueB, ok := value2.(map[string]interface{})
		if ok {
			return EqualStringKeyInterfaceMaps(valueA, valueB)
		}
	case []bool:
		valueB := value2.([]bool)
		return EqualBoolSlices(valueA, valueB)
	case []float64:
		valueB := value2.([]float64)
		return EqualFloat64Slices(valueA, valueB)
	case []string:
		valueB := value2.([]string)
		return EqualStringSlices(valueA, valueB)
	case []map[string]interface{}:
		// In the context of JSON, interface{} must be one of the types of values from json.Unmarshal. We take care of map[string]interface{} specificly here.
		valueB, ok := value2.([]map[string]interface{})
		if ok {
			return EqualStringKeyInterfaceMapSlices(valueA, valueB)
		}
	case reflect.Kind:
		valueB, ok := value2.(reflect.Kind)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case map[string]reflect.Kind:
		valueB, ok := value2.(map[string]reflect.Kind)
		if ok {
			value1A := map[string]interface{}{}
			value2B := map[string]interface{}{}
			for s1, x1 := range valueA {
				value1A[s1] = interface{}(x1)
			}
			for s2, x2 := range valueB {
				value2B[s2] = interface{}(x2)
			}
			return EqualInterfaces(value1A, value2B)
		}
	default:
		fmt.Printf("ERROR EqualValues: Type:%+v Value:%+v for key is out of current options.\n", reflect.TypeOf(value1), reflect.ValueOf(value1))
	}
	return false
}

// EqualStringKeyInterfaceMapSlices finds out if two slices are identical.
func EqualStringKeyInterfaceMapSlices(slice1 []map[string]interface{}, slice2 []map[string]interface{}) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if !EqualStringKeyInterfaceMaps(x, slice2[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// EqualBoolSlices finds out if two slices of strings are identical.
func EqualBoolSlices(slice1 []bool, slice2 []bool) bool {
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

// EqualFloat64Slices finds out if two slices of strings are identical.
func EqualFloat64Slices(slice1 []float64, slice2 []float64) bool {
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

// EqualStringKeyInterfaceMaps finds out if two maps with key typed in string are identical.
func EqualStringKeyInterfaceMaps(map1 map[string]interface{}, map2 map[string]interface{}) bool {
	if len(map1) == len(map2) {
		for key1, value1 := range map1 {
			if !EqualInterfaces(value1, map2[key1]) {
				return false
			}
		}
		return true
	}
	return false
}
