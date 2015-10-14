package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

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
			t.Errorf("ERR TestTypeForValue: interface{}: %+v should be %+v. Not %+v\n", x.In, x.Out, k)
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
			t.Errorf("ERR LowestReflectValue: interface{}: %+v should be %+v. Not %+v\n", x.In, x.Out, v)
		}
	}
	if pass {
		fmt.Println("PASS TestLowestReflectValue")
	}
}
