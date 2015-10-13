package jsonUtil

import (
	"fmt"
	"testing"
)

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
