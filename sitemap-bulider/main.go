package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url" // NEW
	"os"
	"strings" // NEW

	"golang.org/x/net/html"
)

type Link struct {
	Href string
}

type URL struct {
	Loc string `xml:"loc"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
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

	visited := make(map[string]bool)
	crawl("https://example.com", visited)
	generateSitemap(visited)
}
func dfs(h *html.Node, l* []Link){
	if h.Type == html.ElementNode && h.Data == "a" {
	  for _, attr := range h.Attr {
		  if attr.Key == "href" {
			  href := attr.Val
			  fmt.Println(href)
			  *l = append(*l, Link{Href: href})
		  }
	  }
	}
	for c := h.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, l)
	}
}
func crawl(urlString string, visited map[string]bool) {
	if visited[urlString] {
		return
	}

	visited[urlString] = true

	fmt.Println("Visiting:", urlString)

	resp,err := http.Get(urlString)
	if err != nil {
		fmt.Println("Error Fetching:", urlString)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing:", urlString)
		return
	}

	var links []Link
	dfs(doc, &links)

	// Parse current page URL once
	baseURL, err := url.Parse(urlString)
	if err != nil {
		return
	}

	for _, link := range links {

		// Skip invalid links
		if strings.HasPrefix(link.Href, "#") ||
			strings.HasPrefix(link.Href, "mailto:") ||
			strings.HasPrefix(link.Href, "javascript:") {
			continue
		}

		// Convert href into URL object
		u, err := url.Parse(link.Href)
		if err != nil {
			continue
		}

		// Convert relative -> absolute
		absoluteURL := baseURL.ResolveReference(u)

		// Stay on same domain only
		if absoluteURL.Host != baseURL.Host {
			continue
		}

		crawl(absoluteURL.String(), visited)
	}
}
func generateSitemap(visited map[string]bool){
	var urlSet URLSet
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

	for page := range visited {
		urlSet.URLs = append(urlSet.URLs, URL{
			Loc: page,
		})
	}

	data, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return
	}

	data = append([]byte(xml.Header), data...)

	err = os.WriteFile("sitemap.xml", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}