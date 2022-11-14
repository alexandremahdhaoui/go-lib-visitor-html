package visitor

import (
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Edge struct {
	start api.Node
	end   api.Node
}

func (e *Edge) Start() api.Node { return e.start }
func (e *Edge) End() api.Node   { return e.end }

func NewNode(node *html.Node) Node { return Node{node: *node, adjComputed: false, edgeComputed: false} }

type Node struct {
	node  html.Node
	adj   []api.Node
	edges []api.Edge

	adjComputed  bool
	edgeComputed bool
}

func (n *Node) Type() html.NodeType    { return n.node.Type }
func (n *Node) DataAtom() atom.Atom    { return n.node.DataAtom }
func (n *Node) Data() string           { return n.node.Data }
func (n *Node) Namespace() string      { return n.node.Namespace }
func (n *Node) Attr() []html.Attribute { return n.node.Attr }

func (n *Node) AdjacentNodes() []api.Node {
	if n.adjComputed {
		return n.adj
	}

	var nodes []api.Node
	for c := n.node.FirstChild; n.node.FirstChild != nil; c = c.NextSibling {
		if c == nil {
			break
		}
		nodes = append(nodes, &Node{
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

func (n *Node) Edges() []api.Edge {
	if n.edgeComputed {
		return n.edges
	}
	var edges []api.Edge
	for _, adj := range n.AdjacentNodes() {
		edges = append(edges, &Edge{
			start: n,
			end:   adj,
		})
	}
	n.edges = edges
	n.edgeComputed = true
	return edges
}
