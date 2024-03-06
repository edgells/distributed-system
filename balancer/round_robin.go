package balancer

import "sync/atomic"

type Node struct {
	name string
}

type RoundRobinBalancer[T any] struct {
	nodes   []*T
	total   int
	current int64
}

func NewRoundRobinBalancer[T any](nodes []*T) *RoundRobinBalancer[T] {
	return &RoundRobinBalancer[T]{nodes: nodes, total: len(nodes)}
}

func (rr *RoundRobinBalancer[T]) NextNode() *T {
	if len(rr.nodes) == 0 {
		return nil
	}

	node := rr.nodes[rr.current]
	atomic.AddInt64(&rr.current, 1)
	return node
}
