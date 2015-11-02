package jsonUtil

import "reflect"

// ValueTypesInMap is a MapValueType slice to store all the types for values in a map.
type ValueTypesInMap []MapValueType

// func (v ValueTypesInMap) GetNodeValue(m MapValueType) ValueTypesInMap {
// 	for k, v := range m {
// 		out, found := v.Find(parentKey, parentMapValueType)
//
// 	}
// }

// A node in the type map is represented by MapValueType.
// Linked node is the node that has at least one of ParentKey/ClosestKey is the same as the target node's ParentKey/ClosestKey.
// Chain is a path composed with multiple linked nodes with no sub-branch.

// func (v ValueTypesInMap) FindLowerBoundInLinkedNodes(m MapValueType) ValueTypesInMap {
// 	r := ValueTypesInMap{}
// 	for for _, vv := range v {
// 		if m.ClosestKey == vv.ParentKey || m.ClosestKey == vv.ClosestKey && m.NoOfSliceLevels >= 0 {
// 			r = append(r, vv)
// 		}
// 	}
// }
//
//
//
// // These are used as outcome of CompareDepthOfTwoNodes.
// const (
// 	FirstDeeper int = iota
// 	SecondDeeper
// 	Same
// )
//
// // CompareDepthOfTwoNodes compares two nodes from the same ValueTypesInMap to find out which one is deeper. Only the nodes on the same chain is comparable
// func (v ValueTypesInMap)CompareDepthOfTwoNodes(first MapValueType, second MapValueType) (out int, comparable bool) {
//
// }
//
// func (v ValueTypesInMap) AllChains() []ValueTypesInMap {
//
// }
//
// // A top node is a node has no parent node and smallest value in NoOfSliceLevels.
// func (v ValueTypesInMap) TopNodes() []ValueTypesInMap {
// 	r := []ValueTypesInMap{}
// 	if len(v) > 0
// 	for _, x := range v {
// 		i := 0
// 		for _, y := range v {
// 			if y.ClosestKey == x.ParentKey {
// 				i++
// 				o, found := v.FindForKeys(x.ParentKey, x.CloestKey)
// 				if found {
// 					if len(o) > 1 {
//
// 					}
// 				}
// 				break
// 			}
// 		}
// 		if i > 0 {
// 			r = append(r, x)
// 		}
// 	}
//
// }
//
// // FindTopBottomNodes returns the top and bottom nodes. If the receiver is not a chain, it will return error.
// func (v ValueTypesInMap) FindTopBottomNodes(parentKey string, key string) (top MapValueType, bottom MapValueType, err error) {
//
// }

// Chains turns a slice of ValueTypesInMap, which represents a chain that follows the sequence from top to bottom.
// Steps to get all chains:
// 1. Find out the upper and lower linked nodes for each node in receiver
// 2. Find out the top nodes among receiver
// 3. Start chain exploration from top nodes
// func (v ValueTypesInMap) Chains() (chains []ValueTypesInMap) {
// 	r := ValueTypesInMap{}
// 	a := AllNodeWithNeighbor()
// 	t := AllTopNodes(a)
// }

// AllSubchainsForNode gets all subchains for a node.
func (v ValueTypesInMap) AllSubchainsForNode(node MapValueType) []ValueTypesInMap {
	c := BaseChain(v)
	c.Chains = []ValueTypesInMap{ValueTypesInMap{node}}
	expanded := true
	for expanded {
		var h []ValueTypesInMap
		h, expanded = c.ExpandToOneMoreLevelSubchain()
		if expanded {
			c.Chains = h
		}
	}
	return c.Chains
}

// AllNodeWithNeighbor returns NodeWithNeighbor for each node in receiver.
func (v ValueTypesInMap) AllNodeWithNeighbor() []NodeWithNeighbor {
	n := []NodeWithNeighbor{}
	for _, x := range v {
		neighbor, _ := v.NodeWithNeighbor(x)
		n = append(n, neighbor)
	}
	return n
}

// NodeWithNeighbor returns the node with its upper and lower nodes. If the node is not in the receiver, it returns nodeInReceiver = false.
func (v ValueTypesInMap) NodeWithNeighbor(node MapValueType) (out NodeWithNeighbor, nodeInReceiver bool) {
	r := NodeWithNeighbor{}
	_, found := v.Find(node.ParentKey, node.ClosestKey, node.ParentNoOfSliceLevels, node.NoOfSliceLevels)
	if found {
		r.Node = node
		r.Upper = v.FindOneLevelUpperNodesInChain(node)
		r.Lower = v.FindOneLevelLowerNodesInChain(node)
		return r, true
	}
	return r, false
}

// FindOneLevelUpperNodesInChain returns all the nodes that are one level higher.
func (v ValueTypesInMap) FindOneLevelUpperNodesInChain(m MapValueType) ValueTypesInMap {
	r := ValueTypesInMap{}
	for _, vv := range v {
		if vv.ClosestKey == m.ParentKey && vv.NoOfSliceLevels == m.ParentNoOfSliceLevels && m.NoOfSliceLevels == 0 || m.ClosestKey == vv.ClosestKey && vv.NoOfSliceLevels == m.NoOfSliceLevels+1 {
			r = append(r, vv)
		}
	}
	return r
}

// FindOneLevelLowerNodesInChain returns all the nodes that are one level lower.
func (v ValueTypesInMap) FindOneLevelLowerNodesInChain(m MapValueType) ValueTypesInMap {
	r := ValueTypesInMap{}
	for _, vv := range v {
		if m.ClosestKey == vv.ParentKey && m.NoOfSliceLevels == vv.ParentNoOfSliceLevels && vv.NoOfSliceLevels == 0 || m.ClosestKey == vv.ClosestKey && m.NoOfSliceLevels == vv.NoOfSliceLevels+1 {
			r = append(r, vv)
		}
	}
	return r
}

// FindForKeys finds MapValueType that only satisfy having the same keys.
func (v ValueTypesInMap) FindForKeys(parentKey string, key string) (out []MapValueType, found bool) {
	r := []MapValueType{}
	for _, x := range v {
		if parentKey == x.ParentKey && key == x.ClosestKey {
			r = append(r, x)
		}
	}
	if len(r) > 0 {
		return r, true
	}
	return r, false
}

// Find returns the MapValueType that matches the search criteria. It returns false, if no match found.
func (v ValueTypesInMap) Find(parentKey string, key string, parentNoOfSliceLevels int, noOfSliceLvs int) (out MapValueType, found bool) {
	for _, x := range v {
		if x.IsForSameNode(MapValueType{parentKey, key, parentNoOfSliceLevels, noOfSliceLvs, reflect.Invalid}) {
			return x, true
		}
	}
	return MapValueType{}, false
}

// Delete removes a MapValueType from the slice.
func (v ValueTypesInMap) Delete(parentKey string, key string, parentNoOfSliceLevels int, noOfSliceLvs int) ValueTypesInMap {
	vv := ValueTypesInMap{}
	for _, x := range v {
		if !x.IsForSameNode(MapValueType{parentKey, key, parentNoOfSliceLevels, noOfSliceLvs, reflect.Invalid}) {
			vv = append(vv, x)
		}
	}
	return vv
}

// Append appends a MapValueType to ValueTypesInMap. It treats the schema of the JSON as fixed and let other mechanism to handle dynamic JSON schema, if needed. A new slice is always returned. the only difference is the appended value that indicates if the value is really appended.
func (v ValueTypesInMap) Append(m MapValueType) (out ValueTypesInMap, appended bool) {
	o, found := v.Find(m.ParentKey, m.ClosestKey, m.ParentNoOfSliceLevels, m.NoOfSliceLevels)
	if found {
		if o.Type == reflect.Invalid {
			if m.Type == reflect.Invalid {
				return v, false
			}
			v = v.Delete(m.ParentKey, m.ClosestKey, m.ParentNoOfSliceLevels, m.NoOfSliceLevels)
		} else {
			return v, false
		}
	}
	v = append(v, m)
	return v, true
}
