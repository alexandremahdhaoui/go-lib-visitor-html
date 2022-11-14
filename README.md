# go-lib-visitor-html

## Usage

```golang
func printData(n *html.Node) {
    fmt.Println(n.Data)
}

func main() {
	...
    doc, err := html.Parse(bytes.NewReader(b))
    go_lib_html_visitor.VisitNode(doc, printData)
}
```