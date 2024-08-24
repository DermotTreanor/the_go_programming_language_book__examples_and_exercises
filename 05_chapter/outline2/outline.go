// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body) //html.Parse converts the body into a recursive structure. Returns the root of that structure.
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement) //pass the funcs whose behaviour we want to run to a func that runs them on each node
	//!-call

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	//We run the whole func on all children and siblings first. So, by the time we run post on any one node, its sibs and childs are done
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode || n.Type == html.CommentNode{
		//Indent 2 spaces per depth before displaying data.
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data) //The %*s takes in a minimum width and pads with spaces to get that.
		//The actual string we substitute is empty, "". So it will always be padded.
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
