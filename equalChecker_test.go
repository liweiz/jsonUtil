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
var f1 = []float64{1, 2, 3}
var f2 = []float64{1}
var f3 = []float64{2, 3, 1}
var f4 = []float64{1, 4, 3}
var b1 = []bool{true, false, true}
var b2 = []bool{true}
var b3 = []bool{false, true, true}
var b4 = []bool{true, true, true}
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

type InOutEqualInterfaces struct {
	TestTitle string
	Ins       []interface{}
	Out       bool
}

func TestEqualInterfaces(t *testing.T) {
	fmt.Println("\nSTART TestEqualInterfaces")
	toTest := []InOutEqualInterfaces{
		InOutEqualInterfaces{
			"Nil and map",
			[]interface{}{nil, m9},
			false,
		},
		InOutEqualInterfaces{
			"Same bools",
			[]interface{}{true, true},
			true,
		},
		InOutEqualInterfaces{
			"Different bools",
			[]interface{}{true, false},
			false,
		},
		InOutEqualInterfaces{
			"Same float64s",
			[]interface{}{0.1, 0.1},
			true,
		},
		InOutEqualInterfaces{
			"Different float64s",
			[]interface{}{0.1, 0.2},
			false,
		},
		InOutEqualInterfaces{
			"Same strings",
			[]interface{}{"m9", "m9"},
			true,
		},
		InOutEqualInterfaces{
			"Different strings",
			[]interface{}{"m9", ""},
			false,
		},
		InOutEqualInterfaces{
			"Same interface maps",
			[]interface{}{m9, m9},
			true,
		},
		InOutEqualInterfaces{
			"Different interface maps",
			[]interface{}{m9, m1},
			false,
		},
		InOutEqualInterfaces{
			"Same bool slices",
			[]interface{}{b1, b1},
			true,
		},
		InOutEqualInterfaces{
			"Different bool slices",
			[]interface{}{b1, b2},
			false,
		},
		InOutEqualInterfaces{
			"Same float64 slices",
			[]interface{}{f1, f1},
			true,
		},
		InOutEqualInterfaces{
			"Different float64 slices",
			[]interface{}{f1, f2},
			false,
		},
		InOutEqualInterfaces{
			"Same string slices",
			[]interface{}{s1, s1},
			true,
		},
		InOutEqualInterfaces{
			"Different string slices",
			[]interface{}{s1, s2},
			false,
		},
		InOutEqualInterfaces{
			"Same interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1, m3}},
			true,
		},
		InOutEqualInterfaces{
			"Different interface map slices",
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1}},
			false,
		},

		InOutEqualInterfaces{
			"Same reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.Bool},
			true,
		},
		InOutEqualInterfaces{
			"Different reflect.Kinds",
			[]interface{}{reflect.Bool, reflect.String},
			false,
		},
		InOutEqualInterfaces{
			"Same reflect.Kind maps",
			[]interface{}{m13, m13},
			true,
		},
		InOutEqualInterfaces{
			"Different reflect.Kind maps",
			[]interface{}{m13, m14},
			false,
		},
		InOutEqualInterfaces{
			"Same not supported slices",
			[]interface{}{[]int{8, 9}, []int{8, 9}},
			false,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualInterfaces(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualInterfaces: %+v interface{}: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualInterfaces: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualInterfaces")
	}
}

type InOutEqualStringKeyInterfaceMapSlices struct {
	TestTitle string
	Ins       [][]map[string]interface{}
	Out       bool
}

func TestEqualStringKeyInterfaceMapSlices(t *testing.T) {
	fmt.Println("\nSTART TestEqualStringKeyInterfaceMapSlices")
	toTest := []InOutEqualStringKeyInterfaceMapSlices{
		InOutEqualStringKeyInterfaceMapSlices{
			"Same map slices",
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m1, m10}},
			true,
		},
		InOutEqualStringKeyInterfaceMapSlices{
			"Different map slices with different elements",
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m10}},
			false,
		},
		InOutEqualStringKeyInterfaceMapSlices{
			"Different map slices with same keys but different values",
			[][]map[string]interface{}{[]map[string]interface{}{m4}, []map[string]interface{}{m9}},
			false,
		},
		InOutEqualStringKeyInterfaceMapSlices{
			"Different map slices with one empty",
			[][]map[string]interface{}{[]map[string]interface{}{}, []map[string]interface{}{m10}},
			false,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualStringKeyInterfaceMapSlices(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualStringKeyInterfaceMapSlices: %+v Slices: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualStringKeyInterfaceMapSlices: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualStringKeyInterfaceMapSlices")
	}
}

type InOutEqualBoolSlices struct {
	TestTitle string
	Ins       [][]bool
	Out       bool
}

func TestEqualBoolSlices(t *testing.T) {
	fmt.Println("\nSTART TestEqualBoolSlices")
	toTest := []InOutEqualBoolSlices{
		InOutEqualBoolSlices{
			"Different slices with different numbers of elements",
			[][]bool{b1, b2},
			false,
		},
		InOutEqualBoolSlices{
			"Different slices with different orders of elements",
			[][]bool{b1, b3},
			false,
		},
		InOutEqualBoolSlices{
			"Different slices with different elements",
			[][]bool{b1, b4},
			false,
		},
		InOutEqualBoolSlices{
			"Same slices",
			[][]bool{b1, b1},
			true,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualBoolSlices(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualBoolSlices: %+v: Slices: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualBoolSlices: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualBoolSlices")
	}
}

type InOutEqualFloat64Slices struct {
	TestTitle string
	Ins       [][]float64
	Out       bool
}

func TestEqualFloat64Slices(t *testing.T) {
	fmt.Println("\nSTART TestEqualFloat64Slices")
	toTest := []InOutEqualFloat64Slices{
		InOutEqualFloat64Slices{
			"Different slices with different numbers of elements",
			[][]float64{f1, f2},
			false,
		},
		InOutEqualFloat64Slices{
			"Different slices with different orders of elements",
			[][]float64{f1, f3},
			false,
		},
		InOutEqualFloat64Slices{
			"Different slices with different elements",
			[][]float64{f1, f4},
			false,
		},
		InOutEqualFloat64Slices{
			"Same slices",
			[][]float64{f1, f1},
			true,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualFloat64Slices(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualFloat64Slices: %+v: Slices: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualFloat64Slices: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualFloat64Slices")
	}
}

type InOutEqualStringSlices struct {
	TestTitle string
	Ins       [][]string
	Out       bool
}

func TestEqualStringSlices(t *testing.T) {
	fmt.Println("\nSTART TestEqualStringSlices")
	toTest := []InOutEqualStringSlices{
		InOutEqualStringSlices{
			"Different slices with different numbers of elements",
			[][]string{s1, s2},
			false,
		},
		InOutEqualStringSlices{
			"Different slices with different orders of elements",
			[][]string{s1, s3},
			false,
		},
		InOutEqualStringSlices{
			"Different slices with different elements",
			[][]string{s1, s4},
			false,
		},
		InOutEqualStringSlices{
			"Same slices",
			[][]string{s1, s1},
			true,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualStringSlices(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualInterfaces: %+v: Slices: %+v AND %+v equal failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualStringSlices: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualStringSlices")
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
		InOutEqualStringKeyInterfaceMaps{
			"Different maps, different keys different nested values",
			[]map[string]interface{}{m10, m12},
			false,
		},
	}
	pass := true
	for _, x := range toTest {
		if EqualStringKeyInterfaceMaps(x.Ins[0], x.Ins[1]) != x.Out {
			pass = false
			t.Errorf("ERR TestEqualStringKeyInterfaceMaps  %+v: StringKeyMaps: %+v AND %+v comparison failed. Should be %+v\n", x.TestTitle, x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("OK TestEqualStringKeyInterfaceMaps: ", x.TestTitle, x.Ins[0], x.Ins[1])
		}
	}
	if pass {
		fmt.Println("PASS TestEqualStringKeyInterfaceMaps")
	}
}
