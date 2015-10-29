package jsonUtil

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
