package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"
	"os"
)

var port = flag.Int("port", 9898, "Port to listen on")

func main() {
	flag.Parse()
	
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting webserver at http://%s:%d\n", host, *port)

	http.HandleFunc("/", webRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func webRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}