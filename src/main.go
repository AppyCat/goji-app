package main

import (
	"bufio"
	"fmt"
	"github.com/unrolled/render"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"os"
	"strings"
)

func main() {
	render := render.New(render.Options{
		Directory:  "src/views",
		Extensions: []string{".html"},
	})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)

	goji.Get("/hello/:name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
	})

	goji.Get("/wow", func(c web.C, w http.ResponseWriter, r *http.Request) {
		render.HTML(w, http.StatusOK, "index", nil)
	})

	goji.Get("/bar", func(c web.C, w http.ResponseWriter, r *http.Request) {
		render.HTML(w, http.StatusOK, "bar", nil)
	})

	goji.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
		render.HTML(w, http.StatusOK, "index", nil)
	})

	goji.Serve()
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "src/assets" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/js"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
