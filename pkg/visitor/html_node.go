package visitor

import (
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type htmlEdge struct {
	start api.Node
	end   api.Node
}

func (e *htmlEdge) Start() api.Node { return e.start }
func (e *htmlEdge) End() api.Node   { return e.end }

type htmlNode struct {
	node  html.Node
	adj   []api.Node
	edges []api.Edge

	adjComputed  bool
	edgeComputed bool
}

func (n *htmlNode) Type() html.NodeType    { return n.node.Type }
func (n *htmlNode) DataAtom() atom.Atom    { return n.node.DataAtom }
func (n *htmlNode) Data() string           { return n.node.Data }
func (n *htmlNode) Namespace() string      { return n.node.Namespace }
func (n *htmlNode) Attr() []html.Attribute { return n.node.Attr }

func (n *htmlNode) AdjacentNodes() []api.Node {
	if n.adjComputed {
		return n.adj
	}
	var nodes []api.Node
	for c := n.node.FirstChild; c.NextSibling != nil; c = c.NextSibling {
		nodes = append(nodes, &htmlNode{
			node:         *c,
			adj:          nil,
			edges:        nil,
			adjComputed:  false,
			edgeComputed: false,
		})
	}
	n.adj = nodes
	n.adjComputed = true
	return nodes
}

func (n *htmlNode) Edges() []api.Edge {
	if n.edgeComputed {
		return n.edges
	}
	var edges []api.Edge
	for _, adj := range n.AdjacentNodes() {
		edges = append(edges, &htmlEdge{
			start: n,
			end:   adj,
		})
	}
	n.edges = edges
	n.edgeComputed = true
	return edges
}
