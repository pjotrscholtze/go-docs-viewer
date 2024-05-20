package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/config"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/controller"
)

func getCliArgumentWithDefaultValue(idx uint, defaultValue string) string {
	if idx >= uint(len(os.Args)) {
		return defaultValue
	}
	return os.Args[idx]
}
func showHelp() {
	println("Usage: docsviewer CMD Args...")
	println("CMD:")
	println("  help   Shows this help screen")
	println("  serve  Starts serving the Markdown")
	println("")
	println("Args for serve:")
	println("  path        The path to serve, default './'")
	println("  listenAddr  The addr to listen on, default ':8080'")
	println("")
	println(strings.Join([]string{
		"Example for serving a folder called docs in the current directory,",
		" on port 80"}, ""))
	println("  docsviewer serve ./docs :80")
	println("")
	println("Example to serve with default values:")
	println("  docsviewer serve")
}

func main() {
	kind := getCliArgumentWithDefaultValue(1, "help")
	if kind == "serve" {
		baseDir, err := filepath.Abs(getCliArgumentWithDefaultValue(2, "./"))
		if err != nil {
			log.Fatalln(err)
		}
		c := config.Config{
			BaseDir: baseDir,
			// BaseDir:    getCliArgumentWithDefaultValue(2, "/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/"),
			ListenAddr: getCliArgumentWithDefaultValue(3, ":8080"),
		}
		http.Handle("/static", http.FileServer(http.Dir("public")))
		http.HandleFunc("/", controller.GenerateMarkdownFunc(c))
		http.ListenAndServe(c.ListenAddr, nil)
		return
	}

	showHelp()
}
