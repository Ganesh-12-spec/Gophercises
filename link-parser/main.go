package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct{
	Href string
	Text string
}



func main() {
	r := strings.NewReader(`<html><body><a href="/about">About</a></body></html>`)
	doc, err := html.Parse(r)

	if err != nil {
		fmt.Println("not able to parse the html")
		return
	}

	var links []Link
	dfs(doc, &links)
	fmt.Println(links)

}

// dfs traverses the HTML tree depth-first to find <a> elements and extract links.
func dfs(h *html.Node, l *[]Link) {
	if h.Type == html.ElementNode && h.Data == "a" {
		var href string
		for _, attribute := range h.Attr {
			if attribute.Key == "href" {
				href = attribute.Val
			}
		}
		text := getText(h)
		*l = append(*l, Link{Href: href, Text: text})
	}
	for c := h.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, l)
	}
}

// getText recursively collects and returns the trimmed text content of a node.
func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	var result string
	for c:= n.FirstChild; c!= nil; c= c.NextSibling{
		result += getText(c)
	}
	return strings.TrimSpace(result)
}
