package controller

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/converter"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/filetree"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/view"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/blockelements"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/parser"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/spanelements"
)

var parseOrder = []func(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement{
	blockelements.ParseLineCodeblockElement,
	blockelements.ParseLineHeaderElement,
	blockelements.ParseParagraphElement,
	blockelements.ParseLineTableElement,
	blockelements.ParseLineHorizontalLineElement,
	blockelements.ParseLineBlockquoteElement,
	blockelements.ParseListContainerElement,
	blockelements.ParseListContainerNestedElement,
	blockelements.ParseLineTermDefinitionLineElement,
	blockelements.ParseFootnoteDefinitionElement,

	spanelements.ParseInlineCodeElement,
	spanelements.ParseLineImageElement,
	spanelements.ParseLineLinkElement,

	spanelements.ParseLineCheckboxElement,

	spanelements.ParseLineBoldElement,
	spanelements.ParseLineBoldAltElement,
	spanelements.ParseLineEmojiElement,
	spanelements.ParseLineFootnoteElement,
	spanelements.ParseLineHighlightElement,
	spanelements.ParseLineItalicElement,
	spanelements.ParseLineItalicAltElement,
	spanelements.ParseLineStrikethroughElement,
}

func ListOfFiles(basePath string, files []filetree.FileEntry) htmlwrapper.Elm {
	out := []htmlwrapper.Elm{}
	for _, file := range files {
		icon := "fa-regular fa-file"
		if file.IsDir {
			icon = "fa-regular fa-folder"
		}
		p := file.Path[len(basePath):]
		out = append(out, &htmlwrapper.HTMLElm{
			Tag: "li",
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag: "a",
					Attrs: map[string]string{
						"href": "/" + p,
					},
					Contents: []htmlwrapper.Elm{
						&htmlwrapper.HTMLElm{
							Tag: "i",
							Attrs: map[string]string{
								"class": icon,
							},
						},
						htmlwrapper.Text(" " + p),
					},
				},
			},
		})
	}

	return &htmlwrapper.HTMLElm{
		Tag:      "ul",
		Contents: out,
	}
}

func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	currentFile := "/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/" + r.RequestURI
	if _, err := os.Stat(currentFile); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		view.NotFound(w)
		return
	} else if filetree.IsDir(currentFile) {
		// Is dir.
		files2 := filetree.ScanDirs("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/" + r.RequestURI + "/*")
		files := filetree.ScanDirs("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/*")
		filesList := ListOfFiles("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/", files2)
		title, elm := view.Page(
			r.RequestURI,
			"/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/",
			files,
			filesList,
		)
		html, _ := elm.AsHTML()
		view.Wrap(w, "Directory overview: "+title, html)
		return
	}

	if filepath.Ext(currentFile) != ".md" {
		http.FileServer(http.Dir("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/")).ServeHTTP(w, r)
		return
	}

	content, err := filetree.ReadFile(currentFile)
	if err != nil {
		view.Error(err.Error(), w)
	}
	doc := parser.ParseString(content, parseOrder)
	conv := converter.NewConverter()
	contents := conv.Convert(doc)
	// contents := markdowntohtml.Convert(doc)

	title, elm := view.Page(
		r.RequestURI,
		"/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/",
		filetree.ScanDirs("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/*"),
		&htmlwrapper.MultiElm{
			Contents: contents,
		},
	)
	html, _ := elm.AsHTML()
	view.Wrap(w, title, html)
}
