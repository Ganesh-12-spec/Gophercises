package main

import (
	"fmt"
	"net/http"
)

func main() {
	pathToUrls := map[string]string{
		"/google": "https://www.google.com",
		"/github": "https://www.github.com",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fmt.Fprintln(w, "Path not found")
		}
	})

	http.ListenAndServe(":8080", nil)
}