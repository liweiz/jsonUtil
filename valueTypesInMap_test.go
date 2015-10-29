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
	MapValueType{"", "A", 0, 0, reflect.Invalid},
	MapValueType{"", "B", 0, 0, reflect.Float64},
	MapValueType{"", "C", 0, 0, reflect.String},
	MapValueType{"", "D", 0, 0, reflect.Bool},
	MapValueType{"", "E", 0, 0, reflect.Slice},
	MapValueType{"", "F", 0, 0, reflect.Map},
	MapValueType{"E", "a", 1, 1, reflect.Bool},
	MapValueType{"E", "b", 1, 1, reflect.Invalid},
	MapValueType{"F", "c", 0, 0, reflect.String},
	MapValueType{"F", "d", 0, 0, reflect.Invalid},
}

type InFindValueTypesInMap struct {
	ParentKey          string
	Key                string
	ParentNoOfSliceLvs int
	NoOfSliceLvs       int
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
			InFindValueTypesInMap{"", "B", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"", "B", 0, 0, reflect.Float64},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null root existing path with no slice level",
			InFindValueTypesInMap{"", "A", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"", "A", 0, 0, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null non-root existing path with no slice level",
			InFindValueTypesInMap{"F", "c", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "c", 0, 0, reflect.String},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with no slice level",
			InFindValueTypesInMap{"F", "d", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "d", 0, 0, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null non-root existing path with slice level",
			InFindValueTypesInMap{"E", "a", 1, 1},
			OutFindValueTypesInMap{
				MapValueType{"E", "a", 1, 1, reflect.Bool},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with slice level",
			InFindValueTypesInMap{"E", "b", 1, 1},
			OutFindValueTypesInMap{
				MapValueType{"E", "b", 1, 1, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null non-root existing path with slice level, not able to be found",
			InFindValueTypesInMap{"E", "b", 2, 2},
			OutFindValueTypesInMap{
				MapValueType{},
				false,
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := samlpeValueTypesInMap
		out, found := s.Find(x.Ins.ParentKey, x.Ins.Key, x.Ins.ParentNoOfSliceLvs, x.Ins.NoOfSliceLvs)
		if found == x.Outs.Found {
			if found {
				if out.IsForSameNode(x.Outs.MapValueType) && out.Type == x.Outs.MapValueType.Type {
					fmt.Println("OK TestFindValueTypesInMap: ", x.TestTitle, x.Ins, x.Outs)
				} else {
					pass = false
					t.Errorf("\nERR TestFindValueTypesInMap: %+v find for: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Outs)
				}
			} else {
				fmt.Println("OK TestFindValueTypesInMap: ", x.TestTitle, x.Ins, x.Outs)
			}
		} else {
			pass = false
			t.Errorf("\nERR TestFindValueTypesInMap: %+v find for: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Outs)
		}
	}
	if pass {
		fmt.Println("PASS TestFindValueTypesInMap")
	}
}

type InOutDeleteValueTypesInMap struct {
	TestTitle string
	Ins       InFindValueTypesInMap
	Out       ValueTypesInMap
}

func TestDeleteValueTypesInMap(t *testing.T) {
	fmt.Println("\nSTART TestDeleteValueTypesInMap")
	toTest := []InOutDeleteValueTypesInMap{
		InOutDeleteValueTypesInMap{
			"Delete a non-null root existing path with no slice level",
			InFindValueTypesInMap{"", "B", 0, 0},
			ValueTypesInMap{
				MapValueType{"", "A", 0, 0, reflect.Invalid},
				// MapValueType{"", "B", 0, 0, reflect.Float64},
				MapValueType{"", "C", 0, 0, reflect.String},
				MapValueType{"", "D", 0, 0, reflect.Bool},
				MapValueType{"", "E", 0, 0, reflect.Slice},
				MapValueType{"", "F", 0, 0, reflect.Map},
				MapValueType{"E", "a", 1, 1, reflect.Bool},
				MapValueType{"E", "b", 1, 1, reflect.Invalid},
				MapValueType{"F", "c", 0, 0, reflect.String},
				MapValueType{"F", "d", 0, 0, reflect.Invalid},
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := samlpeValueTypesInMap
		s = s.Delete(x.Ins.ParentKey, x.Ins.Key, x.Ins.ParentNoOfSliceLvs, x.Ins.NoOfSliceLvs)
		ok := true
		if len(s) == len(x.Out) {
			i := 0
			for _, xx := range s {
				for _, xxx := range x.Out {
					if xx.IsForSameNode(xxx) && xx.Type == xxx.Type {
						i++
					}
				}
			}
			if i == len(s) {
				fmt.Println("OK TestDeleteValueTypesInMap: ", x.TestTitle, x.Ins, x.Out)
			} else {
				ok = false
			}
		} else {
			ok = false
		}
		if !ok {
			pass = false
			t.Errorf("\nERR TestDeleteValueTypesInMap: %+v delete for: %+v failed. Should be %+v\n", x.TestTitle, x.Ins, x.Out)
		}
	}
	if pass {
		fmt.Println("PASS TestDeleteValueTypesInMap")
	}
}

type OutAppendValueTypesInMap struct {
	Out      ValueTypesInMap
	Appended bool
}

type InOutAppendValueTypesInMap struct {
	TestTitle string
	In        MapValueType
	Outs      OutAppendValueTypesInMap
}

func TestAppendValueTypesInMap(t *testing.T) {
	fmt.Println("\nSTART TestAppendValueTypesInMap")
	toTest := []InOutAppendValueTypesInMap{
		InOutAppendValueTypesInMap{
			"Append a non-existing path",
			MapValueType{"", "X", 0, 0, reflect.Invalid},
			OutAppendValueTypesInMap{
				ValueTypesInMap{
					MapValueType{"", "A", 0, 0, reflect.Invalid},
					MapValueType{"", "B", 0, 0, reflect.Float64},
					MapValueType{"", "C", 0, 0, reflect.String},
					MapValueType{"", "D", 0, 0, reflect.Bool},
					MapValueType{"", "E", 0, 0, reflect.Slice},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"E", "a", 1, 1, reflect.Bool},
					MapValueType{"E", "b", 1, 1, reflect.Invalid},
					MapValueType{"F", "c", 0, 0, reflect.String},
					MapValueType{"F", "d", 0, 0, reflect.Invalid},
					MapValueType{"", "X", 0, 0, reflect.Invalid},
				},
				true,
			},
		},
		InOutAppendValueTypesInMap{
			"Append an existing path with invalid type",
			MapValueType{"", "A", 0, 0, reflect.Invalid},
			OutAppendValueTypesInMap{
				ValueTypesInMap{
					MapValueType{"", "A", 0, 0, reflect.Invalid},
					MapValueType{"", "B", 0, 0, reflect.Float64},
					MapValueType{"", "C", 0, 0, reflect.String},
					MapValueType{"", "D", 0, 0, reflect.Bool},
					MapValueType{"", "E", 0, 0, reflect.Slice},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"E", "a", 1, 1, reflect.Bool},
					MapValueType{"E", "b", 1, 1, reflect.Invalid},
					MapValueType{"F", "c", 0, 0, reflect.String},
					MapValueType{"F", "d", 0, 0, reflect.Invalid},
				},
				false,
			},
		},
		InOutAppendValueTypesInMap{
			"Append an existing path with valid type",
			MapValueType{"", "A", 0, 0, reflect.Bool},
			OutAppendValueTypesInMap{
				ValueTypesInMap{
					MapValueType{"", "B", 0, 0, reflect.Float64},
					MapValueType{"", "C", 0, 0, reflect.String},
					MapValueType{"", "D", 0, 0, reflect.Bool},
					MapValueType{"", "E", 0, 0, reflect.Slice},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"E", "a", 1, 1, reflect.Bool},
					MapValueType{"E", "b", 1, 1, reflect.Invalid},
					MapValueType{"F", "c", 0, 0, reflect.String},
					MapValueType{"F", "d", 0, 0, reflect.Invalid},
					MapValueType{"", "A", 0, 0, reflect.Invalid},
				},
				true,
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := samlpeValueTypesInMap
		ss, added := s.Append(x.In)
		ok := true
		if len(ss) == len(x.Outs.Out) {
			if added == x.Outs.Appended {
				n := 0
				for i, xx := range ss {
					if xx.IsForSameNode(x.Outs.Out[i]) && xx.Type == x.Outs.Out[i].Type {
						n++
					}
				}
				if n == len(ss) {
					fmt.Println("OK TestAppendValueTypesInMap: ", x.TestTitle, x.In, x.Outs)
				} else {
					ok = false
				}
			} else {
				ok = false
			}
		} else {
			ok = false
		}
		if !ok {
			pass = false
			t.Errorf("\nERR TestAppendValueTypesInMap: %+v append for: %+v failed. Should be %+v\n", x.TestTitle, x.In, x.Outs)
		}
	}
	if pass {
		fmt.Println("PASS TestAppendValueTypesInMap")
	}
}
