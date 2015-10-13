package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

var s1 = []string{"a", "b", "c"}
var s2 = []string{"a"}
var s3 = []string{"b", "c", "a"}
var s4 = []string{"a", "d", "c"}
var s5 = []string{"a", "b", "c"}
var m1 = map[string]interface{}{"string": "a"}
var m2 = map[string]interface{}{"bool": true}
var m3 = map[string]interface{}{"float64": float64(45)}
var m4 = map[string]interface{}{"[]string": []string{"a", "b", "c"}}
var m5 = map[string]interface{}{"map": m1}
var m6 = map[string]interface{}{"string": "b"}
var m7 = map[string]interface{}{"bool": false}
var m8 = map[string]interface{}{"float64": float64(43)}
var m9 = map[string]interface{}{"[]string": []string{"d", "b", "c"}}
var m10 = map[string]interface{}{"map": m6}
var m11 = map[string]interface{}{"bool": true}
var m12 = map[string]interface{}{"bool": true, "map": m6}
var m13 = map[string]reflect.Kind{"bool": reflect.Bool, "map": reflect.Map}
var m14 = map[string]reflect.Kind{"invalid": reflect.Invalid}

// type InOutFindTypeForValue struct {
// 	In  interface{}
// 	Out reflect.Kind
// }
//
// func TestFindTypeForValue(t *testing.T) {
// 	fmt.Println("\nSTART TestFindTypeForValue")
// 	toTest := []InOutFindTypeForValue{
// 		// String
// 		InOutFindTypeForValue{"a", reflect.String},
// 		// Bool
// 		InOutFindTypeForValue{false, reflect.Bool},
// 		// Float64
// 		InOutFindTypeForValue{4.1, reflect.Float64},
// 		// Map
// 		InOutFindTypeForValue{map[string]interface{}{"a": 2}, reflect.Map},
// 		// Slice
// 		InOutFindTypeForValue{[]string{"b"}, reflect.Slice},
// 		// String reflect
// 		InOutFindTypeForValue{reflect.ValueOf("a"), reflect.String},
// 		// Bool reflect
// 		InOutFindTypeForValue{reflect.ValueOf(false), reflect.Bool},
// 		// Float64 reflect
// 		InOutFindTypeForValue{reflect.ValueOf(4.1), reflect.Float64},
// 		// Map reflect
// 		InOutFindTypeForValue{reflect.ValueOf(map[string]interface{}{"a": 2}), reflect.Map},
// 		// Slice reflect
// 		InOutFindTypeForValue{reflect.ValueOf([]string{"b"}), reflect.Slice},
// 	}
// 	for _, x := range toTest {
// 		k := FindTypeForValue(x.In)
// 		if k == x.Out {
// 			fmt.Println("TestFindTypeForValue passed: ", k)
// 		} else {
// 			t.Errorf("ERR: interface{}: %+v TestFindTypeForValue should be %+v. Not %+v\n", x.In, x.Out, k)
// 		}
// 	}
// }
//
// type InOutTestLowestReflectValue struct {
// 	In  interface{}
// 	Out reflect.Value
// }
//
// func TestLowestReflectValue(t *testing.T) {
// 	fmt.Println("\nSTART TestLowestReflectValue")
// 	toTest := []InOutTestLowestReflectValue{
// 		// Non reflect.Value
// 		InOutTestLowestReflectValue{"a", reflect.ValueOf("a")},
// 		// 1 level reflect.Value
// 		InOutTestLowestReflectValue{reflect.ValueOf("a"), reflect.ValueOf("a")},
// 		// 5 level reflect.Value
// 		InOutTestLowestReflectValue{reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf("a"))))), reflect.ValueOf("a")},
// 	}
// 	for _, x := range toTest {
// 		v := LowestReflectValue(x.In)
// 		if v.String() == x.Out.String() {
// 			fmt.Println("TestLowestReflectValue passed: ", v)
// 		} else {
// 			t.Errorf("ERR: interface{}: %+v LowestReflectValue should be %+v. Not %+v\n", x.In, x.Out, v)
// 		}
// 	}
// }

type InOutTestEqualValues struct {
	TestTitle string
	Ins       []interface{}
	Out       bool
}

func TestEqualValues(t *testing.T) {
	fmt.Println("\nSTART TestEqualValues")
	toTest := []InOutTestEqualValues{
		InOutTestEqualValues{
			"Nil and map",
			[]interface{}{nil, m9},
			false,
		},
		InOutTestEqualValues{
			"Same bools",
			[]interface{}{true, true},
			true,
		},
		InOutTestEqualValues{
			"Different bools",
			[]interface{}{true, false},
			false,
		},
		InOutTestEqualValues{
			"Same float64s",
			[]interface{}{0.1, 0.1},
			true,
		},
		InOutTestEqualValues{
			"Different float64s",
			[]interface{}{0.1, 0.2},
			false,
		},
		InOutTestEqualValues{
			"Same strings",
			[]interface{}{"m9", "m9"},
			true,
		},
		InOutTestEqualValues{
			"Different strings",
			[]interface{}{"m9", ""},
			false,
		},
		InOutTestEqualValues{
			"Same interface maps",
			[]interface{}{m9, m9},
			true,
		},
		InOutTestEqualValues{
			"Different interface maps",
			[]interface{}{m9, m1},
			false,
		},
		InOutTestEqualValues{
			"Same string slices",
			[]interface{}{s1, s1},
			true,
		},
		InOutTestEqualValues{
			"Different string slices",
			[]interface{}{s1, s2},
			false,
		},
		InOutTestEqualValues{
			"Same interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1, m3}},
			true,
		},
		InOutTestEqualValues{
			"Different interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1}},
			false,
		},

		InOutTestEqualValues{
			"Same reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.Bool},
			true,
		},
		InOutTestEqualValues{
			"Different reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.String},
			false,
		},
		InOutTestEqualValues{
			"Same reflect.Kind maps",
			[]interface{}{m13, m13},
			true,
		},
		InOutTestEqualValues{
			"Different reflect.Kind maps",
			[]interface{}{m13, m14},
			false,
		},
		InOutTestEqualValues{
			"Same not supported slices",
			[]interface{}{[]float64{0.8, 0.9}, []float64{0.8, 0.9}},
			false,
		},
	}
	for _, x := range toTest {
		if EqualValues(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR TestEqualValues: %+v interface{}: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("PASS TestEqualValues: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
}
