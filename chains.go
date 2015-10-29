package jsonUtil

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
func (c Chains) ExpandToOneMoreLevelSubchain() (chains []ValueTypesInMap, expanded bool) {
	expanded = false
	chains = []ValueTypesInMap{}
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
					chains = append(chains, y)
					expanded = true
				}
			}
		} else {
			chainExpandable = true
		}
		if !chainExpandable {
			chains = append(chains, ch)
		}
	}
	return chains, expanded
}
