package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

type InOutValueTypeInMap struct {
	TestTitle string
	In        map[string]interface{}
	Out       map[string]reflect.Kind
}

// func TestValueTypeInMap(t *testing.T) {
// 	fmt.Println("\nSTART TestValueTypeInMap")
// 	toTest := []InOutValueTypeInMap{
// 		InOutValueTypeInMap{"All values valid map", m4, map[string]reflect.Kind{"[]string": reflect.Slice}},
// 		InOutValueTypeInMap{"Map with some invalid value", m15, map[string]reflect.Kind{"bool": reflect.Bool, "map": reflect.Map, "i": reflect.Invalid}},
// 	}
// 	pass := true
// 	for _, x := range toTest {
// 		if !EqualInterfaces(ValueTypeInMap(x.In), x.Out) {
// 			pass = false
// 			t.Errorf("ERR TestValueTypeInMap: %+v failed. Should be %+v\n", x.In, x.Out)
// 		} else {
// 			fmt.Println("OK TestValueTypeInMap: ", x.TestTitle, x.In)
// 		}
// 	}
// 	if pass {
// 		fmt.Println("PASS TestValueTypeInMap")
// 	}
// }

type InOutTypeForValue struct {
	TestTitle string
	In        interface{}
	Out       reflect.Kind
}

func TestTypeForValue(t *testing.T) {
	fmt.Println("\nSTART TestTypeForValue")
	toTest := []InOutTypeForValue{
		InOutTypeForValue{"String", "a", reflect.String},
		InOutTypeForValue{"Bool", false, reflect.Bool},
		InOutTypeForValue{"Float64", 4.1, reflect.Float64},
		InOutTypeForValue{"Map", map[string]interface{}{"a": 2}, reflect.Map},
		InOutTypeForValue{"Slice", []interface{}{"b"}, reflect.Slice},
		InOutTypeForValue{"Nil", nil, reflect.Invalid},
		InOutTypeForValue{"String reflect", reflect.ValueOf("a"), reflect.String},
		InOutTypeForValue{"Bool reflect", reflect.ValueOf(false), reflect.Bool},
		InOutTypeForValue{"Float64 reflect", reflect.ValueOf(4.1), reflect.Float64},
		InOutTypeForValue{"Map reflect", reflect.ValueOf(map[string]interface{}{"a": 2}), reflect.Map},
		InOutTypeForValue{"Slice reflect", reflect.ValueOf([]interface{}{"b"}), reflect.Slice},
		// InOutTypeForValue{"Interface String reflect", reflect.ValueOf(interface{}("a")), reflect.String},
		// InOutTypeForValue{"Interface Bool reflect", reflect.ValueOf(interface{}(false)), reflect.Bool},
		// InOutTypeForValue{"Interface Float64 reflect", reflect.ValueOf(interface{}(4.1)), reflect.Float64},
		// InOutTypeForValue{"Interface Map reflect", reflect.ValueOf(interface{}(map[string]interface{}{"a": 2})), reflect.Map},
		// InOutTypeForValue{"Interface Slice reflect", reflect.ValueOf(interface{}([]interface{}{"b"})), reflect.Slice},
	}
	pass := true
	for _, x := range toTest {
		k := TypeForValue(x.In)
		if k == x.Out {
			fmt.Println("OK TestTypeForValue: ", x.TestTitle, k)
		} else {
			pass = false
			t.Errorf("\nERR TestTypeForValue: interface{}: %+v should be %+v. Not %+v\n", x.In, x.Out, k)
		}
	}
	if pass {
		fmt.Println("PASS TestTypeForValue")
	}
}

type InOutLowestReflectValue struct {
	TestTitle string
	In        interface{}
	Out       reflect.Value
}

func TestLowestReflectValue(t *testing.T) {
	fmt.Println("\nSTART TestLowestReflectValue")
	toTest := []InOutLowestReflectValue{
		InOutLowestReflectValue{"Non reflect.Value", "a", reflect.ValueOf("a")},
		InOutLowestReflectValue{"1 level reflect.Value", reflect.ValueOf("a"), reflect.ValueOf("a")},
		InOutLowestReflectValue{"5 level reflect.Value", reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf("a"))))), reflect.ValueOf("a")},
	}
	pass := true
	for _, x := range toTest {
		v := LowestReflectValue(x.In)
		if v.String() == x.Out.String() {
			fmt.Println("OK TestLowestReflectValue: ", x.TestTitle, v)
		} else {
			pass = false
			t.Errorf("\nERR LowestReflectValue: interface{}: %+v should be %+v. Not %+v\n", x.In, x.Out, v)
		}
	}
	if pass {
		fmt.Println("PASS TestLowestReflectValue")
	}
}
