package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/carlmjohnson/gateway"
)

type Link struct {
	Title       string
	Description string
	Image       string
	Link        string
}

var links map[string][]Link = map[string][]Link{
	"Links": {
		{Title: "SOURCE CODE", Description: "contibute to bungolo", Image: "github-mark-white.png", Link: "https://github.com/bungolo-dev"},
	},
}

func main() {

	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()

	home := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=300")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, links)
	}

 img := func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)

	}

	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", 8080)
		listener = http.ListenAndServe
	}

 http.HandleFunc("/", home)
 http.HandleFunc("/api/img/*", img)
	log.Fatal(listener(portStr, nil))
}
