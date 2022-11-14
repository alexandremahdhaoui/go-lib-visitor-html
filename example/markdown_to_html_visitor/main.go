package main

import (
	"bytes"
	"github.com/yuin/goldmark"
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/graph"
	"gitlab.com/alexandre.mahdhaoui/go-lib-visitor-html/pkg/visitor"
	"golang.org/x/net/html"
	"io"
	"os"
)

func displayTextNodeData(node *visitor.Node) bool { println(node.Data()); return false }

func main() {
	doc := getHtmlDoc()

	// initialize the node
	node := visitor.NewNode(doc)

	// build the visitor
	builder := visitor.NewBuilder()
	v := builder.SetVisitTextNode(displayTextNodeData).Build()

	// Visit the root node
	v.Visit(&node)

	// Visit node's graph with depth-first search algorithm:
	graph.DFS(&node, &v)

	// Visit node's graph with breadth-first search algorithm:
	graph.DFS(&node, &v)
}

func getHtmlDoc() *html.Node {
	file, err := os.Open("example/markdown_to_html_visitor/raw_markdown.md")
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(b, &buf); err != nil {
		panic(err)
	}

	b, err = io.ReadAll(&buf)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	return doc
}
