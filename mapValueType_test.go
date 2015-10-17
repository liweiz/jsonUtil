package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

type InValueTypeForPath struct {
	ParentKey       string
	ClosestKey      string
	NoOfSliceLevels int
	Interface       interface{}
}

type InOutValueTypeForPath struct {
	TestTitle string
	Ins       InValueTypeForPath
	Out       MapValueType
}

func TestValueTypeForPath(t *testing.T) {
	fmt.Println("\nSTART TestValueTypeForPath")
	toTest := []InOutValueTypeForPath{
		InOutValueTypeForPath{
			"Float64",
			InValueTypeForPath{"", "", 0, 1.1},
			MapValueType{"", "", 0, reflect.Float64},
		},
		InOutValueTypeForPath{
			"Invalid",
			InValueTypeForPath{"", "", 0, nil},
			MapValueType{"", "", 0, reflect.Invalid},
		},
	}
	pass := true
	for _, x := range toTest {
		r := ValueTypeForPath(x.Ins.ParentKey, x.Ins.ClosestKey, x.Ins.NoOfSliceLevels, x.Ins.Interface)
		if r.ClosestKey == x.Out.ClosestKey && r.ParentKey == x.Out.ParentKey && r.NoOfSliceLevels == x.Out.NoOfSliceLevels && r.Type == x.Out.Type {
			fmt.Println("OK TestValueTypeForPath: ", x.TestTitle, x.Ins)
		} else {
			pass = false
			t.Errorf("ERR TestValueTypeForPath: %+v input: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Out)
		}
	}
	if pass {
		fmt.Println("PASS TestValueTypeForPath")
	}
}
