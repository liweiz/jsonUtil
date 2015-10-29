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
var nodeA = MapValueType{
	"",
	"A",
	0,
	0,
	reflect.Invalid,
}

type InOutAllTopNodes struct {
	TestTitle string
	In        []NodeWithNeighbor
	Out       ValueTypesInMap
}

func TestAllTopNodes(t *testing.T) {
	fmt.Println("\nSTART TestAllTopNodes")
	toTest := []InOutAllTopNodes{
		InOutAllTopNodes{
			"Empty []NodeWithNeighbor",
			[]NodeWithNeighbor{},
			ValueTypesInMap{},
		},
		InOutAllTopNodes{
			"One level []NodeWithNeighbor",
			[]NodeWithNeighbor{
				NodeWithNeighbor{
					MapValueType{
						"",
						"A",
						0,
						0,
						reflect.Bool,
					},
					ValueTypesInMap{},
					ValueTypesInMap{},
				},
			},
			ValueTypesInMap{},
		},
		InOutAllTopNodes{
			"Three level []NodeWithNeighbor",
			[]NodeWithNeighbor{
				NodeWithNeighbor{
					MapValueType{
						"",
						"A",
						0,
						0,
						reflect.Map,
					},
					ValueTypesInMap{},
					ValueTypesInMap{},
				},
				NodeWithNeighbor{
					MapValueType{
						"A",
						"B",
						0,
						0,
						reflect.Slice,
					},
					ValueTypesInMap{},
					ValueTypesInMap{},
				},
			},
			ValueTypesInMap{},
		},
	}
}
