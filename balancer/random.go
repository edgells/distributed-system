package balancer

import "math/rand"

type RandomNode struct {
	name string
}

type RandomBalancer struct {
	nodes []*RandomNode
	total int
}

func NewRandomBalancer(nodes []*RandomNode) *RandomBalancer {
	return &RandomBalancer{nodes: nodes, total: len(nodes)}
}

func (r *RandomBalancer) NextNode() *RandomNode {
	if len(r.nodes) == 0 {
		return nil
	}

	random := rand.Intn(r.total + 1)
	return r.nodes[random]
}
