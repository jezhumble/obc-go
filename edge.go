package obc

type HopStrategy func(*edge, *Node, []*Node) int

type edge struct {
	child *Node
	cost int
}

func Edge(child *Node, cost int) *edge {
	edge := new(edge)
	edge.child = child
	edge.cost = cost
	return edge
}

func HopsToStrategy(s *edge, destination *Node, visited []*Node) int {
	return s.child.costTo(destination, visited, HopsToStrategy) + 1
}

func CostToStrategy(s *edge, destination *Node, visited []*Node) int {
	return s.child.costTo(destination, visited, CostToStrategy) + s.cost
}
