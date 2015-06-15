package obc

import (
	"testing"
)

var (
	a, b, c, d, e, f, g, h = new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node)
)

func TestMain(t *testing.T) {
	h.AddChild(Edge(b, 2))
	b.AddChild(Edge(a, 4))
	b.AddChild(Edge(c, 3))
	a.AddChild(Edge(f, 5))
	c.AddChild(Edge(d, 2))
	c.AddChild(Edge(e, 1))
	c.AddChild(Edge(e, 3))
	d.AddChild(Edge(e, 5))
	e.AddChild(Edge(b, 1))
}
	
func TestCanReachSelf(t *testing.T) {
	node := new(Node)
	if !node.CanReach(node) {
		t.Fatalf("Node couldn't reach itself")
	}
}

func TestCanReachChild(t *testing.T) {
	node := new (Node)
	child := new (Node)
	node.AddChild(Edge(child, 1))
	if !node.CanReach(child) {
		t.Fatalf("Node couldn't reach its child")
	}
}

func TestCanReachDescendents(t *testing.T) {
	if (!h.CanReach(e)) {
		t.Fatalf("Can't reach descendent nodes")
	}
}

func TestCanReachChildWithCycles(t *testing.T) {
	if h.CanReach(g) {
		t.Fatalf("Shouldn't be able to reach disconnected node")
	}
}

func TestHopCount(t *testing.T) {
	hops := h.HopsTo(c)
	if hops != 2 {
		t.Fatalf("C should be 2 hops from H but was %d", hops)
	}
}

func TestMinHopCount(t *testing.T) {
	hops := h.HopsTo(e)
	if hops != 3 {
		t.Fatalf("E should be 3 hops from H but was %d", hops)
	}
}

func TestCostBetweenNodes(t *testing.T) {
	cost := h.CostTo(c)
	if cost != 5 {
		t.Fatalf("H to C should cost 5 but was %d", cost)
	}
}
