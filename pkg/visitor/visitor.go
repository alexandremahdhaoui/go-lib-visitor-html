// Package visitor
// TODO:
//   - use a Struct to properly Visit the tree:
//     https://www.lihaoyi.com/post/ZeroOverheadTreeProcessingwiththeVisitorPattern.html
//   - In the Visitor struct add `strategy` field to allow different Tree traversal algorithm.
//     https://en.wikipedia.org/wiki/Tree_traversal
//   - Maybe create a proper `go_lib_visitor_tree` package to handle Tree traversal.
package visitor

import (
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
	"golang.org/x/net/html"
)

type Visitor interface {
	api.Visitor
	VisitErrorNode(htmlNode) bool
	VisitTextNode(htmlNode) bool
	VisitDocumentNode(htmlNode) bool
	VisitElementNode(htmlNode) bool
	VisitCommentNode(htmlNode) bool
	VisitDoctypeNode(htmlNode) bool
}

func Visit(node htmlNode, visitor Visitor) bool {
	switch node.node.Type {
	case html.ErrorNode:
		return visitor.VisitErrorNode(node)
	case html.TextNode:
		return visitor.VisitTextNode(node)
	case html.DocumentNode:
		return visitor.VisitDocumentNode(node)
	case html.ElementNode:
		return visitor.VisitElementNode(node)
	case html.CommentNode:
		return visitor.VisitCommentNode(node)
	case html.DoctypeNode:
		return visitor.VisitDoctypeNode(node)
	default:
		panic("not implemented")
	}
}
