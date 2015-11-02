package jsonUtil

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"
// )
//
// var SampleNodes = map[string]NodeWithNeighbor{
// 	"A": NodeWithNeighbor{
// 		SamlpeValueTypesInMap[0],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{},
// 	},
// 	"B": NodeWithNeighbor{
// 		SamlpeValueTypesInMap[1],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{},
// 	},
// 	"C": NodeWithNeighbor{
// 		SamlpeValueTypesInMap[2],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{},
// 	},
// 	"D": NodeWithNeighbor{
// 		SamlpeValueTypesInMap[3],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{},
// 	},
// 	"E": NodeWithNeighbor{
// 		SamlpeValueTypesInMap[4],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{
// 			MapValueType{"E", "a", 0, 1, reflect.Bool},
// 			MapValueType{"E", "b", 0, 1, reflect.Invalid},
// 		},
// 	},
// 	NodeWithNeighbor{
// 		SamlpeValueTypesInMap[5],
// 		ValueTypesInMap{},
// 		ValueTypesInMap{
// 			MapValueType{"F", "c", 0, 0, reflect.String},
// 			MapValueType{"F", "d", 0, 0, reflect.Invalid},
// 		},
// 	},
// 	NodeWithNeighbor{
// 		SamlpeValueTypesInMap[6],
// 		ValueTypesInMap{MapValueType{"", "E", 0, 0, reflect.Slice}},
// 		ValueTypesInMap{},
// 	},
// 	NodeWithNeighbor{
// 		SamlpeValueTypesInMap[7],
// 		ValueTypesInMap{MapValueType{"", "E", 0, 0, reflect.Slice}},
// 		ValueTypesInMap{},
// 	},
// 	NodeWithNeighbor{
// 		SamlpeValueTypesInMap[8],
// 		ValueTypesInMap{MapValueType{"", "F", 0, 0, reflect.Map}},
// 		ValueTypesInMap{},
// 	},
// 	NodeWithNeighbor{
// 		SamlpeValueTypesInMap[9],
// 		ValueTypesInMap{MapValueType{"", "F", 0, 0, reflect.Map}},
// 		ValueTypesInMap{},
// 	},
// }
//
// type InOutAllTopNodes struct {
// 	TestTitle string
// 	In        []NodeWithNeighbor
// 	Out       ValueTypesInMap
// }
//
// func TestAllTopNodes(t *testing.T) {
// 	fmt.Println("\nSTART TestAllTopNodes")
// 	toTest := []InOutAllTopNodes{
// 		InOutAllTopNodes{
// 			"Empty []NodeWithNeighbor",
// 			[]NodeWithNeighbor{},
// 			ValueTypesInMap{},
// 		},
// 		InOutAllTopNodes{
// 			"One level []NodeWithNeighbor",
// 			[]NodeWithNeighbor{},
// 			ValueTypesInMap{},
// 		},
// 		InOutAllTopNodes{
// 			"Three level []NodeWithNeighbor",
// 			[]NodeWithNeighbor{
// 				NodeWithNeighbor{
// 					MapValueType{
// 						"",
// 						"A",
// 						0,
// 						0,
// 						reflect.Map,
// 					},
// 					ValueTypesInMap{},
// 					ValueTypesInMap{},
// 				},
// 				NodeWithNeighbor{
// 					MapValueType{
// 						"A",
// 						"B",
// 						0,
// 						0,
// 						reflect.Slice,
// 					},
// 					ValueTypesInMap{},
// 					ValueTypesInMap{},
// 				},
// 			},
// 			ValueTypesInMap{},
// 		},
// 	}
// }
