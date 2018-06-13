package main

import (
	"fmt"
	"log"
	"net/http"
	debug "github.com/AlexisVisco/debug"
)

var httpdeb, _ = debug.Register("http")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	httpdeb.Log(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
}

func main() {

	// manually set option of httpdeb for portability
	httpdeb.Option.Color = true
	httpdeb.Option.Enabled = true
	httpdeb.Option.Latency = true

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
