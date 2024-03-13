package balancer

import (
	"math/rand"
	"sync"
)

type WrrNode struct {
	node   string
	weight int
	sync.Map
}

type WrrBalance struct {
	nodes       []*WrrNode
	totalWeight int
}

func NewWrrBalance(nodes []*WrrNode) *WrrBalance {
	var totalWeight int
	for _, node := range nodes {
		totalWeight += node.weight
	}

	return &WrrBalance{nodes: nodes, totalWeight: totalWeight}
}

func (b *WrrBalance) NextNode() *WrrNode {
	if len(b.nodes) == 0 {
		return nil
	}

	var curWeight int
	randWeight := b.randomWeight()
	for _, node := range b.nodes {
		curWeight += node.weight
		if curWeight > randWeight {
			return node
		}
	}

	return b.nodes[0]
}

// random weight interface
func (b *WrrBalance) randomWeight() int {
	return rand.Intn(b.totalWeight + 1)
}
