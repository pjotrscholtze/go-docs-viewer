package main

import (
	"net/http"

	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/config"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/controller"
)

func main() {
	c := config.Config{
		ListenAddr: ":8080",
		BaseDir:    "/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/",
	}
	http.Handle("/static", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", controller.GenerateMarkdownFunc(c))
	http.ListenAndServe(c.ListenAddr, nil)
}
