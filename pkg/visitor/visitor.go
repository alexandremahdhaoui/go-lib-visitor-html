package visitor

import (
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
	"golang.org/x/net/html"
)

type (
	Func      func(node *Node) bool
	Interface interface {
		VisitCommentNode(*Node) bool
		VisitDoctypeNode(*Node) bool
		VisitDocumentNode(*Node) bool
		VisitElementNode(*Node) bool
		VisitErrorNode(*Node) bool
		VisitTextNode(*Node) bool
	}
)

func NewFunc() Func { return func(node *Node) bool { return false } }

// New instantiate a new Visitor performing no actions when visiting a node.
// Please use the Builder to
func New() Visitor {
	return Visitor{
		visitCommentNode:  NewFunc(),
		visitDoctypeNode:  NewFunc(),
		visitDocumentNode: NewFunc(),
		visitElementNode:  NewFunc(),
		visitErrorNode:    NewFunc(),
		visitTextNode:     NewFunc(),
	}
}

type Visitor struct {
	visitCommentNode  Func
	visitDoctypeNode  Func
	visitDocumentNode Func
	visitElementNode  Func
	visitErrorNode    Func
	visitTextNode     Func
}

func (v *Visitor) Visit(n api.Node) bool {
	node := n.(*Node)
	switch node.node.Type {
	case html.CommentNode:
		return v.VisitCommentNode(node)
	case html.DoctypeNode:
		return v.VisitDoctypeNode(node)
	case html.DocumentNode:
		return v.VisitDocumentNode(node)
	case html.ElementNode:
		return v.VisitElementNode(node)
	case html.ErrorNode:
		return v.VisitErrorNode(node)
	case html.TextNode:
		return v.VisitTextNode(node)
	default:
		panic("not implemented")
	}
}

func (v *Visitor) VisitCommentNode(node *Node) bool  { return v.visitCommentNode(node) }
func (v *Visitor) VisitDoctypeNode(node *Node) bool  { return v.visitDoctypeNode(node) }
func (v *Visitor) VisitDocumentNode(node *Node) bool { return v.visitDocumentNode(node) }
func (v *Visitor) VisitElementNode(node *Node) bool  { return v.visitElementNode(node) }
func (v *Visitor) VisitErrorNode(node *Node) bool    { return v.visitErrorNode(node) }
func (v *Visitor) VisitTextNode(node *Node) bool     { return v.visitTextNode(node) }
