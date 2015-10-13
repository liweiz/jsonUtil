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

type InOutTestEqualInterfaces struct {
	TestTitle string
	Ins       []interface{}
	Out       bool
}

func TestEqualInterfaces(t *testing.T) {
	fmt.Println("\nSTART TestEqualInterfaces")
	toTest := []InOutTestEqualInterfaces{
		InOutTestEqualInterfaces{
			"Nil and map",
			[]interface{}{nil, m9},
			false,
		},
		InOutTestEqualInterfaces{
			"Same bools",
			[]interface{}{true, true},
			true,
		},
		InOutTestEqualInterfaces{
			"Different bools",
			[]interface{}{true, false},
			false,
		},
		InOutTestEqualInterfaces{
			"Same float64s",
			[]interface{}{0.1, 0.1},
			true,
		},
		InOutTestEqualInterfaces{
			"Different float64s",
			[]interface{}{0.1, 0.2},
			false,
		},
		InOutTestEqualInterfaces{
			"Same strings",
			[]interface{}{"m9", "m9"},
			true,
		},
		InOutTestEqualInterfaces{
			"Different strings",
			[]interface{}{"m9", ""},
			false,
		},
		InOutTestEqualInterfaces{
			"Same interface maps",
			[]interface{}{m9, m9},
			true,
		},
		InOutTestEqualInterfaces{
			"Different interface maps",
			[]interface{}{m9, m1},
			false,
		},
		InOutTestEqualInterfaces{
			"Same string slices",
			[]interface{}{s1, s1},
			true,
		},
		InOutTestEqualInterfaces{
			"Different string slices",
			[]interface{}{s1, s2},
			false,
		},
		InOutTestEqualInterfaces{
			"Same interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1, m3}},
			true,
		},
		InOutTestEqualInterfaces{
			"Different interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1}},
			false,
		},

		InOutTestEqualInterfaces{
			"Same reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.Bool},
			true,
		},
		InOutTestEqualInterfaces{
			"Different reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.String},
			false,
		},
		InOutTestEqualInterfaces{
			"Same reflect.Kind maps",
			[]interface{}{m13, m13},
			true,
		},
		InOutTestEqualInterfaces{
			"Different reflect.Kind maps",
			[]interface{}{m13, m14},
			false,
		},
		InOutTestEqualInterfaces{
			"Same not supported slices",
			[]interface{}{[]float64{0.8, 0.9}, []float64{0.8, 0.9}},
			false,
		},
	}
	for _, x := range toTest {
		if EqualInterfaces(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR TestEqualInterfaces: %+v interface{}: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("PASS TestEqualInterfaces: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
}

type InOutTestEqualStringSlices struct {
	TestTitle string
	Ins       [][]string
	Out       bool
}

type testStringKeyMapSlicesInBoolOut struct {
	Ins [][]map[string]interface{}
	Out bool
}

func TestCompareStringKeyMapSlices(t *testing.T) {
	fmt.Println("\nSTART TestCompareStringKeyMapSlices")
	toTest := []testStringKeyMapSlicesInBoolOut{
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m1, m10}},
			true,
		},
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m10}},
			false,
		},
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{}, []map[string]interface{}{m10}},
			false,
		},
	}
	for _, x := range toTest {
		if CompareStringKeyMapSlices(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: Slices: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}

func TestCompareStringSlices(t *testing.T) {
	fmt.Println("\nSTART TestCompareStringSlices")
	toTest := []testStringSliceInBoolOut{
		testStringSliceInBoolOut{
			[][]string{s1, s2},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s3},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s4},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s5},
			true,
		},
	}
	for _, x := range toTest {
		if CompareStringSlices(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: Slices: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}

type InOutEqualStringKeyInterfaceMaps struct {
	TestTitle string
	Ins       []map[string]interface{}
	Out       bool
}

func TestEqualStringKeyInterfaceMaps(t *testing.T) {
	fmt.Println("\nSTART TestEqualStringKeyInterfaceMaps")
	toTest := []InOutEqualStringKeyInterfaceMaps{
		InOutEqualStringKeyInterfaceMaps{
			"Same maps",
			[]map[string]interface{}{m9, m9},
			true,
		},
		InOutEqualStringKeyInterfaceMaps{
			"Different maps, same keys different values",
			[]map[string]interface{}{m7, m11},
			false,
		},
		InOutEqualStringKeyInterfaceMaps{
			"Different maps, same keys different nested values",
			[]map[string]interface{}{m4, m9},
			false,
		},
		InOutEqualStringKeyInterfaceMaps{
			"Different maps, different keys different nested values",
			[]map[string]interface{}{m10, m12},
			false,
		},
	}
	for _, x := range toTest {
		if EqualStringKeyInterfaceMaps(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR TestEqualStringKeyInterfaceMaps: StringKeyMaps: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("PASS TestEqualStringKeyInterfaceMaps: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
}

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
