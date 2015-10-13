package jsonUtil

// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

// Data decoding flow:
// 1. loaded in as interface{}
// 2. find out type
// 2.1 in known keys
// 3 convert to value
// 2.2 unknown key
// 2.2.1 explore with reflect
// 2.2.1.1 nil/empty, mark reflect.Invalid
// 2.2.1.2 map, mark reflect.Map
// 2.2.1.2.1 cast from interface{} to map[string]interface{}
// 2.2.1.2.2 go through all the key value pairs with step 1
// 2.2.1.3 slice, mark reflect.Slice
// 2.2.1.3.1 cast from interface{} to []interface{}
// 2.2.1.3.2 get the first value and go to step 1
// 2.2.1.3 bool/string/float64, mark reflect.Bool/reflect.String/reflect.Float64

import (
	"fmt"
	"reflect"
)

// EqualValues compares two values' equal. Currently targeted at value in bool/string/float64/reflect.Kind and complex value map/slice derived from them.
func EqualValues(value1 interface{}, value2 interface{}) bool {
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
			var value1A, value2B map[string]interface{}
			for s1, x1 := range valueA {
				value1A[s1] = interface{}(x1)
			}
			for s2, x2 := range valueB {
				value2B[s2] = interface{}(x2)
			}
			return EqualValues(value1A, value2B)
		}
	default:
		fmt.Printf("ERROR CompareValues: Type:%+v Value:%+v for key is out of current options.\n", reflect.TypeOf(value1), reflect.ValueOf(value1))
	}
	return false
}
