package go_lib_html_visitor

import (
	"golang.org/x/net/html"
)

func VisitNode(n *html.Node, f func(n *html.Node)) {
	switch n.Type {
	case html.TextNode:
		print("content: ")

	case html.DocumentNode:
		print("document:")
	case html.ElementNode:
		print("element: ")
	case html.DoctypeNode:
		print("doctype: ")
	}
	f(n)
	if n.FirstChild != nil {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			VisitNode(c, f)
		}
	}
}
