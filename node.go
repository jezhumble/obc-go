package obc

import (
	"math"
)

type Node struct {
	children []*edge
}

var UNREACHABLE int = math.MaxInt32

func (s *Node) AddChild(child *edge) {
	s.children = append(s.children, child)
}

func (s *Node) CanReach(destination *Node) bool {
	return s.costTo(destination, make([]*Node, 10), HopsToStrategy) != UNREACHABLE
}

func (s *Node) CostTo(destination *Node) int {
	return s.costTo(destination, make([]*Node, 10), CostToStrategy)
}

func (s *Node) HopsTo(destination *Node) int {
	return s.costTo(destination, make([]*Node, 10), HopsToStrategy)
}

func (s *Node) costTo(destination *Node, visited []*Node, strategy HopStrategy) int {
	for _, v := range visited { if (v == s) { return UNREACHABLE }}
	if s == destination { return 0 }
	var minHops = UNREACHABLE;
	for i := range s.children {
		hops := strategy(s.children[i], destination, append(visited, s))
		if hops < minHops { minHops = hops }
	}
	return minHops
}
