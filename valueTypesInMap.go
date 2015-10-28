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
func (v ValueTypesInMap) Chains() (chains []ValueTypesInMap) {
	r := ValueTypesInMap{}
	a := AllNodeWithNeighbor()
	t := AllTopNodes(a)
}

// func (v ValueTypesInMap) AllSubchainsForNode(node MapValueType) []ValueTypesInMap {
//
// }

// Chains stores the base and its chains derived from the base.
type Chains struct {
	Chains         []ValueTypesInMap
	Base           ValueTypesInMap
	NodesNeighbors []NodeWithNeighbor
}

// BaseChain returns a Chains with Base and NodesInfo set.
func BaseChain(base ValueTypesInMap) Chains {
	c := Chains{}
	c.Base = base
	c.NodesNeighbors = c.Base.AllNodeWithNeighbor()
	return c
}

// ExpandToOneMoreLevelSubchain adds one more lower level nodes to chains where applicable. Empty chain will not be included.
func (c Chains) ExpandToOneMoreLevelSubchain() []ValueTypesInMap {
	r := []ValueTypesInMap{}
	// Go through all known chains.
	for _, ch := range c.Chains {
		chainExpandable := false
		if len(ch) > 0 {
			// Find NodeWithNeighbor for last node in each known chain.
			n, found := FindInNodeWithNeighbors(c.NodesNeighbors, ch[len(ch)-1])
			if found {
				chainExpandable = true
				for _, x := range n.Lower {
					y := append(ch, x)
					r = append(r, y)
				}
			}
		} else {
			chainExpandable = true
		}
		if !chainExpandable {
			r = append(r, ch)
		}
	}
	return r
}

// NodeWithNeighbor is the node with its upper and lower nodes.
type NodeWithNeighbor struct {
	Node  MapValueType
	Upper ValueTypesInMap
	Lower ValueTypesInMap
}

// FindInNodeWithNeighbors returns the MapValueType that matches the search criteria. It returns false, if no match found.
func FindInNodeWithNeighbors(n []NodeWithNeighbor, v MapValueType) (node NodeWithNeighbor, found bool) {
	for _, x := range n {
		if v.IsForSameNode(x.Node) {
			return x, true
		}
	}
	return NodeWithNeighbor{}, false
}

// AllTopNodes returns all top nodes.
func AllTopNodes(n []NodeWithNeighbor) ValueTypesInMap {
	t := []NodeWithNeighbor{}
	for _, x := range n {
		if len(x.Upper) == 0 {
			t = append(t, x)
		}
	}
	r := ValueTypesInMap{}
	for _, y := range t {
		r = append(r, y.Node)
	}
	return r
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
		if parentKey == x.ParentKey && key == x.ClosestKey && parentNoOfSliceLevels == x.ParentNoOfSliceLevels && noOfSliceLvs == x.NoOfSliceLevels {
			return x, true
		}
	}
	return MapValueType{}, false
}

// Delete removes a MapValueType from the slice.
func (v ValueTypesInMap) Delete(parentKey string, key string, parentNoOfSliceLevels int, noOfSliceLvs int) ValueTypesInMap {
	vv := ValueTypesInMap{}
	for _, x := range v {
		if !(parentKey == x.ParentKey && key == x.ClosestKey && parentNoOfSliceLevels == x.ParentNoOfSliceLevels && noOfSliceLvs == x.NoOfSliceLevels) {
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
