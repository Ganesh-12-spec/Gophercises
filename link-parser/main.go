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



func main (){
 r := strings.NewReader(`<html><body><a
	href="/about">About</a></body></html>`)
 doc ,err := html.Parse(r)

	if err != nil {
		fmt.Println("NOt abel to Parser the html link")
	}else{
		fmt.Println(doc.Type)
	}
}