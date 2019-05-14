package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	host := flag.String("host", "0.0.0.0", "host name to bind to")
	port := flag.Int("port", 8080, "port number")
	flag.Parse()
	http.HandleFunc("/", handler)
	bind := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Starting web server on %s\n", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
