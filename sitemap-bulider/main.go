package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
}

func main() {
	uri := "https://example.com"
	resp, err := http.Get(uri)

	if err != nil{
		fmt.Printf("Error: %v\n",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	doc ,err := html.Parse(resp.Body)

	if err != nil {
		fmt.Println("Not able to Parser the html")
		return
	}
	var links []Link
	dfs (doc,&links)
	fmt.Println(links)
}
func dfs(h *html.Node, l* []Link){
	if h.Type == html.ElementNode && h.Data =="a"{
	  for _, attr := range h.Attr {
		  if attr.Key == "href" {
			  fmt.Println(attr.Val)
		  }
	  }
	}
	for c := h.FirstChild; c != nil; c = c.NextSibling{
		dfs(c,l)
	}
}