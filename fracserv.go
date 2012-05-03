package main

import (
	"flag"
	"fmt"
	"github.com/sergeykv/fracserv/fractal"
	_ "github.com/sergeykv/fracserv/fractal/solid"
	"html/template"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

type encodeFunc func(io.Writer, image.Image) error

var port = flag.Int("port", 9898, "Port to listen on")

func main() {
	flag.Parse()

	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting webserver at http://%s:%d\n", host, *port)

	for _, f := range fractal.Names() {
		http.HandleFunc("/form/"+f, acceptsOnly("/form/"+f, formFunc(f)))
		http.HandleFunc("/render/"+f, acceptsOnly("/render/"+f, renderFunc(f, png.Encode)))
		http.HandleFunc("/png/"+f, acceptsOnly("/png/"+f,
			asFile(fmt.Sprintf("%s.png", f), renderFunc(f, png.Encode))))
		http.HandleFunc("/jpeg/"+f, acceptsOnly("/jpeg/"+f,
			asFile(fmt.Sprintf("%s.jpg", f), renderFunc(f, jpegEncode))))
	}
	http.HandleFunc("/", acceptsOnly("/", index))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func acceptsOnly(url string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != url {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}
}

func jpegEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

func executeTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", name))
	if err == nil {
		err = tmpl.Execute(w, data)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formFunc(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		executeTemplate(w, name, nil)
	}
}

func renderFunc(name string, encode encodeFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, err := fractal.NewFractal(name, r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = encode(w, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func asFile(name string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/force-download")
		w.Header().Set("Content-Disposition", "attachment; filename=" + name)
		handler(w, r)
	}
}

func pngFunc(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "index", fractal.Names())
}
