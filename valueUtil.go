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
