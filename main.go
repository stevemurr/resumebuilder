package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"resumebuilder/handler"
)

var addr = flag.String("addr", ":9001", "port to listen on")
var banner = `
Resumebuilder!
`

func main() {
	flag.Parse()
	h := &handler.Handler{}

	http.HandleFunc("/api/toml/pdf", h.ParseTOML)
	http.HandleFunc("/api/toml/tex", h.GetTexFromTOML)

	fmt.Printf("%s\n", banner)
	log.Printf("Listening on %s\n", *addr)

	http.ListenAndServe(*addr, nil)
}
