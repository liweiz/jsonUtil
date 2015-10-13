package jsonUtil

import (
	"fmt"
	"testing"
)

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
