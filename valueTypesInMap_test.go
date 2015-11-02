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
//     false,
//     true
//   ],
//   "F": {
//     "a": "goo",
//     "b": null
//   },
// 	"G": [
//     [
// 			"aa",
// 			"bb"
// 		]
//   ],
// }
var SamlpeValueTypesInMap = ValueTypesInMap{
	MapValueType{"", "A", 0, 0, reflect.Invalid},
	MapValueType{"", "B", 0, 0, reflect.Float64},
	MapValueType{"", "C", 0, 0, reflect.String},
	MapValueType{"", "D", 0, 0, reflect.Bool},
	MapValueType{"", "E", 0, 0, reflect.Slice},
	MapValueType{"", "E", 0, 1, reflect.Bool},
	MapValueType{"", "F", 0, 0, reflect.Map},
	MapValueType{"F", "a", 0, 0, reflect.String},
	MapValueType{"F", "b", 0, 0, reflect.Invalid},
	MapValueType{"", "G", 0, 0, reflect.Slice},
	MapValueType{"", "G", 0, 2, reflect.Slice},
	MapValueType{"", "G", 1, 2, reflect.String},
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
		// No slice level
		InOutFindValueTypesInMap{
			"Find a non-null direct-key-parent existing path, root key",
			InFindValueTypesInMap{"", "B", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"", "B", 0, 0, reflect.Float64},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null direct-key-parent existing path, root key",
			InFindValueTypesInMap{"", "A", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"", "A", 0, 0, reflect.Invalid},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null direct-key-parent existing path, non-root key",
			InFindValueTypesInMap{"F", "a", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "a", 0, 0, reflect.String},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a null direct-key-parent existing path, non-root key",
			InFindValueTypesInMap{"F", "b", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"F", "b", 0, 0, reflect.Invalid},
				true,
			},
		},
		// With slice level
		InOutFindValueTypesInMap{
			"Find a non-null direct-key-parent existing path, root key",
			InFindValueTypesInMap{"", "E", 0, 0},
			OutFindValueTypesInMap{
				MapValueType{"", "E", 0, 0, reflect.Slice},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a non-null non-direct-key-parent existing path, root key",
			InFindValueTypesInMap{"", "E", 0, 1},
			OutFindValueTypesInMap{
				MapValueType{"", "E", 0, 1, reflect.Bool},
				true,
			},
		},
		InOutFindValueTypesInMap{
			"Find a nested non-null non-direct-key-parent existing path, root key",
			InFindValueTypesInMap{"", "G", 1, 2},
			OutFindValueTypesInMap{
				MapValueType{"", "G", 1, 2, reflect.String},
				true,
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := SamlpeValueTypesInMap
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
				MapValueType{"", "E", 0, 1, reflect.Bool},
				MapValueType{"", "F", 0, 0, reflect.Map},
				MapValueType{"F", "a", 0, 0, reflect.String},
				MapValueType{"F", "b", 0, 0, reflect.Invalid},
				MapValueType{"", "G", 0, 0, reflect.Slice},
				MapValueType{"", "G", 0, 2, reflect.Slice},
				MapValueType{"", "G", 1, 2, reflect.String},
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := SamlpeValueTypesInMap
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
					MapValueType{"", "E", 0, 1, reflect.Bool},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"F", "a", 0, 0, reflect.String},
					MapValueType{"F", "b", 0, 0, reflect.Invalid},
					MapValueType{"", "G", 0, 0, reflect.Slice},
					MapValueType{"", "G", 0, 2, reflect.Slice},
					MapValueType{"", "G", 1, 2, reflect.String},
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
					MapValueType{"", "E", 0, 1, reflect.Bool},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"F", "a", 0, 0, reflect.String},
					MapValueType{"F", "b", 0, 0, reflect.Invalid},
					MapValueType{"", "G", 0, 0, reflect.Slice},
					MapValueType{"", "G", 0, 2, reflect.Slice},
					MapValueType{"", "G", 1, 2, reflect.String},
				},
				false,
			},
		},
		InOutAppendValueTypesInMap{
			"Append an existing path with valid type",
			MapValueType{"", "A", 0, 0, reflect.Bool},
			OutAppendValueTypesInMap{
				ValueTypesInMap{
					// MapValueType{"", "A", 0, 0, reflect.Invalid},
					MapValueType{"", "B", 0, 0, reflect.Float64},
					MapValueType{"", "C", 0, 0, reflect.String},
					MapValueType{"", "D", 0, 0, reflect.Bool},
					MapValueType{"", "E", 0, 0, reflect.Slice},
					MapValueType{"", "E", 0, 1, reflect.Bool},
					MapValueType{"", "F", 0, 0, reflect.Map},
					MapValueType{"F", "a", 0, 0, reflect.String},
					MapValueType{"F", "b", 0, 0, reflect.Invalid},
					MapValueType{"", "G", 0, 0, reflect.Slice},
					MapValueType{"", "G", 0, 2, reflect.Slice},
					MapValueType{"", "G", 1, 2, reflect.String},
					MapValueType{"", "A", 0, 0, reflect.Bool},
				},
				true,
			},
		},
	}
	pass := true
	for _, x := range toTest {
		s := SamlpeValueTypesInMap
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
