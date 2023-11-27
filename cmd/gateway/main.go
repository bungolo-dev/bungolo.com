package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/carlmjohnson/gateway"
)

func main() {

	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	http.Handle("/api/feed", &home{})
	http.Handle("/api/home", &home{})

	log.Fatal(listener(portStr, nil))
}

func cacheControlMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=300")
		h.ServeHTTP(w, r)
	})
}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=300")
	w.Write([]byte("This is my home page"))
}
