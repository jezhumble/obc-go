package obc

import (
	"testing"
)

var (
	a, b, c, d, e, f, g, h = new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node)
)

func TestMain(t *testing.T) {
	h.AddChild(b)
	b.AddChild(a)
	b.AddChild(c)
	a.AddChild(f)
	c.AddChild(d)
	c.AddChild(e)
	c.AddChild(e)
	d.AddChild(e)
	e.AddChild(b)
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
	node.AddChild(child)
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
