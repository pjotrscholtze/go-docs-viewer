package main

import (
	"net/http"

	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/controller"
)

func main() {
	http.Handle("/static", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", controller.GenerateMarkdown)
	http.ListenAndServe(":8080", nil)
}
