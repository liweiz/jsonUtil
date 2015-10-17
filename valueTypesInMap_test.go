package jsonUtil

import (
	"fmt"
	"reflect"
	"testing"
)

// Sample JSON:
// {
//   "A": null,
//   "B": 5.5,
//   "C": "This is",
//   "D": true,
//   "E": [
//     "a": false,
//     "b": null
//   ],
//   "F": {
//     "c": "goo",
//     "d": null
//   }
// }
var samlpeValueTypesInMap = ValueTypesInMap{
	MapValueType{"", "A", 0, reflect.Invalid},
	MapValueType{"", "B", 0, reflect.Float64},
	MapValueType{"", "C", 0, reflect.String},
	MapValueType{"", "D", 0, reflect.Bool},
	MapValueType{"", "E", 0, reflect.Slice},
	MapValueType{"", "F", 0, reflect.Map},
	MapValueType{"E", "a", 1, reflect.Bool},
	MapValueType{"F", "c", 1, reflect.String},
	MapValueType{"F", "d", 1, reflect.Invalid},
}

type InFindValueTypesInMap struct {
	ParentKey    string
	Key          string
	NoOfSliceLvs int
}

type OutFindValueTypesInMap struct {
	MapValueType MapValueType
	Found        bool
}

type InOutFindValueTypesInMap struct {
	TestTitle string
	Ins       InFindValueTypesInMap
	Outs      OutFindValueTypesInMap
}

func TestFindValueTypesInMap(t *testing.T) {
	fmt.Println("\nSTART TestFindValueTypesInMap")
	toTest := []InOutFindValueTypesInMap{
		InOutFindValueTypesInMap{
			"Find a non-null root existing path with no slice level",
			InFindValueTypesInMap{"", "B", 0},
			OutFindValueTypesInMap{
				MapValueType{"", "B", 0, reflect.Float64},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null root existing path with no slice level",
			InFindValueTypesInMap{"", "A", 0},
			OutFindValueTypesInMap{
				MapValueType{"", "A", 0, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null non-root existing path with no slice level",
			InFindValueTypesInMap{"F", "c", 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "c", 0, reflect.String},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with no slice level",
			InFindValueTypesInMap{"F", "d", 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "d", 0, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null non-root existing path with slice level",
			InFindValueTypesInMap{"E", "a", 1},
			OutFindValueTypesInMap{
				MapValueType{"E", "a", 1, reflect.Bool},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with slice level",
			InFindValueTypesInMap{"E", "b", 1},
			OutFindValueTypesInMap{
				MapValueType{"E", "b", 1, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with slice level, not able to be found",
			InFindValueTypesInMap{"E", "b", 2},
			OutFindValueTypesInMap{
				MapValueType{},
				false,
			},
		},
	}
	pass := true
	for _, x := range toTest {
		out, found := samlpeValueTypesInMap.Find(x.Ins.ParentKey, x.Ins.Key, x.Ins.NoOfSliceLvs)
		if found == x.Outs.Found {
			if found {
				if out.ParentKey == x.Outs.MapValueType.ParentKey && out.ClosestKey == x.Outs.MapValueType.ClosestKey && out.NoOfSliceLevels == x.Outs.MapValueType.NoOfSliceLevels && out.Type == x.Outs.MapValueType.Type {
					fmt.Println("OK TestFindValueTypesInMap: ", x.TestTitle, x.Ins, x.Outs)
				} else {
					t.Errorf("ERR TestFindValueTypesInMap: %+v find for: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Outs)
				}
			} else {
				fmt.Println("OK TestFindValueTypesInMap: ", x.TestTitle, x.Ins, x.Outs)
			}
		} else {
			t.Errorf("ERR TestFindValueTypesInMap: %+v find for: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Outs)
		}
	}
	if pass {
		fmt.Println("PASS TestFindValueTypesInMap")
	}
}
