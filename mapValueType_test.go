package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

type InValueTypeForPath struct {
	ParentKey             string
	ClosestKey            string
	ParentNoOfSliceLevels int
	NoOfSliceLevels       int
	Interface             interface{}
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
			InValueTypeForPath{"", "", 0, 0, 1.1},
			MapValueType{"", "", 0, 0, reflect.Float64},
		},
		InOutValueTypeForPath{
			"Invalid",
			InValueTypeForPath{"", "", 0, 0, nil},
			MapValueType{"", "", 0, 0, reflect.Invalid},
		},
	}
	pass := true
	for _, x := range toTest {
		r := ValueTypeForPath(x.Ins.ParentKey, x.Ins.ClosestKey, x.Ins.ParentNoOfSliceLevels, x.Ins.NoOfSliceLevels, x.Ins.Interface)
		if r.IsForSameNode(x.Out) && r.Type == x.Out.Type {
			fmt.Println("OK TestValueTypeForPath: ", x.TestTitle, x.Ins)
		} else {
			pass = false
			t.Errorf("\nERR TestValueTypeForPath: %+v input: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Out)
		}
	}
	if pass {
		fmt.Println("PASS TestValueTypeForPath")
	}
}
