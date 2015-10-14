package jsonUtil

import (
	"fmt"
	"reflect"
)

// ValueTypeInMap explores types for all values in a map.
func ValueTypeInMap(m map[string]interface{}) map[string]reflect.Kind {
	r := map[string]reflect.Kind{}
	for k, v := range m {
		r[k] = TypeForValue(v)
	}
	return r
}

// TypeForValue finds out the type. It handles a data itself and the data after any times of operations in reflect.ValueOf().
func TypeForValue(value interface{}) reflect.Kind {
	v := LowestReflectValue(value)
	switch v.Kind() {
	case reflect.Slice:
		if v.Len() > 0 {
			return reflect.Slice
		}
	case reflect.Map:
		if len(v.MapKeys()) > 0 {
			return reflect.Map
		}
	case reflect.String:
		return reflect.String
	case reflect.Float64:
		return reflect.Float64
	case reflect.Bool:
		return reflect.Bool
	// case reflect.Interface:
	// 	switch v.Interface().(type) {
	// 	case bool:
	// 		return reflect.Bool
	// 	case string:
	// 		return reflect.String
	// 	case float64:
	// 		return reflect.Float64
	// 	case map[string]interface{}:
	// 		return reflect.Map
	// 	case []interface{}:
	// 		return reflect.Slice
	// 	default:
	// 		fmt.Printf("ERR FindTypeForValue: Value:%+v is out of current options. %+v\n", value, reflect.TypeOf(value))
	// 	}
	default:
		fmt.Printf("ERR FindTypeForValue: Value:%+v is out of current options. %+v\n", value, reflect.TypeOf(value))
	}
	return reflect.Invalid
}

// LowestReflectValue gets the closest reflect.Value for an interface{} value.
func LowestReflectValue(value interface{}) reflect.Value {
	v, ok := value.(reflect.Value)
	if ok {
		vv, cool := v.Interface().(reflect.Value)
		if cool {
			return LowestReflectValue(vv)
		}
		return v
	}
	return reflect.ValueOf(value)
}
