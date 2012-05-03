package main

import (
	"image/png"
	"flag"
	"fmt"
	"github.com/sergeykv/fracserv/fractal"
	_ "github.com/sergeykv/fracserv/fractal/solid"
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

	http.HandleFunc("/render", render)
	http.HandleFunc("/", webRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func render(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, err := fractal.NewFractal(r.Form.Get("fractal"), r.Form)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	png.Encode(w, f)
}

func webRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
