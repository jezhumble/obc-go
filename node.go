package obc

import (
	"math"
)

type Node struct {
	children []*Node
}

var UNREACHABLE int = math.MaxInt32

func (s *Node) HopsTo(destination *Node) int {
	return s.hopsTo(destination, make([]*Node, 10))
}

func (s *Node) hopsTo(destination *Node, visited []*Node) int {
	for _, v := range visited { if (v == s) { return UNREACHABLE }}
	if s == destination { return 0 }
	var minHops = UNREACHABLE;
	for i := range s.children {
		hops := s.children[i].hopsTo(destination, append(visited, s))
		if hops < minHops { minHops = hops + 1 }
	}
	return minHops
}

func (s *Node) CanReach(destination *Node) bool {
	return s.hopsTo(destination, make([]*Node, 10)) != UNREACHABLE
}

func (s *Node) AddChild(child *Node) {
	s.children = append(s.children, child)
}
