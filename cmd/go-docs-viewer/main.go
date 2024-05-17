package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/filetree"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/markdowntohtml"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/view"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/blockelements"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/parser"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/spanelements"
)

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
						// <i class="fa-regular fa-file"></i>
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
		Wrap(w, "Directory overview: "+title, html)
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
	doc := parser.ParseString(content, []func(input []entity.MarkdownElement, parserFn func(input string) []entity.MarkdownElement) []entity.MarkdownElement{
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
	})
	_ = doc
	contents := markdowntohtml.Convert(doc)

	title, elm := view.Page(
		r.RequestURI,
		"/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/",
		filetree.ScanDirs("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/*"),
		&htmlwrapper.MultiElm{
			Contents: //[]htmlwrapper.Elm{
			// &htmlwrapper.HTMLElm{
			// 	Tag: "pre",
			// 	Contents: []htmlwrapper.Elm{
			// 		htmlwrapper.Text(content),
			// 	},
			// },
			contents,
			// },
		},
	)
	html, _ := elm.AsHTML()
	Wrap(w, title, html)
}

func main() {
	// ScanDirs("/home/pjotr/go/src/github.com/pjotrscholtze/go-docs-viewer/*")
	http.Handle("/static", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", GenerateMarkdown)
	http.ListenAndServe(":8080", nil)

}

func Wrap(w http.ResponseWriter, title, html string) {
	fmt.Fprint(w, `<!doctype html>
	<html lang="en">
	  <head>
	  <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
	  <meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>`+title+`</title>
	  </head>
	  <body>
		`+html+`
		<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.bundle.min.js"></script>
			<!-- Custom JS -->
		<script>
	$(document).ready(function () {
		$('#sidebarCollapse').on('click', function () {
			$('#sidebar').toggleClass('active');
			$('#content').toggleClass('active');
		});
	});
	
		</script>
		<style>
		body {
			font-family: 'Poppins', sans-serif;
			background: #f0f0f1;
			color:#3c434a;
		}
		.wrapper {
			display: flex;
		}
		#sidebar {
			min-width: 250px;
			max-width: 250px;
			background: #1d2327;
			/* #3c434a; */
			color: #f0f0f1;
			transition: all 0.3s;
		}
		#sidebar li a:hover {
			color:  #72aee6;
		}
		#sidebar li a {
			padding: 0.5em;
			display: block;
			color: #f0f0f1;
		}
		#sidebar li ul {
			padding: 0 0.5em;
		}
		#sidebar.active {
			margin-left: -250px;
		}
		#sidebarToggle {
			transition: all 0.3s;
		}
		#sidebarToggle.active {
			transform: rotate(90deg);
		}
		#content {
			width: 100%;
			padding: 20px;
			min-height: 100vh;
			transition: all 0.3s;
		}
		#content.active {
			/* margin-left: 250px; */
		}
		img {
			display: block;
		}




		</style>

		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
		
		<!-- and it's easy to individually load additional languages -->
		<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/go.min.js"></script>
		
		<script>hljs.highlightAll();</script>
		<style>
		/* Dracula Theme v1.2.5
 *
 * https://github.com/dracula/highlightjs
 *
 * Copyright 2016-present, All rights reserved
 *
 * Code licensed under the MIT license
 *
 * @author Denis Ciccale <dciccale@gmail.com>
 * @author Zeno Rocha <hi@zenorocha.com>
 */

.hljs {
  display: block;
  overflow-x: auto;
  padding: 0.5em;
  background: #282a36;
}

.hljs-built_in,
.hljs-selector-tag,
.hljs-section,
.hljs-link {
  color: #8be9fd;
}

.hljs-keyword {
  color: #ff79c6;
}

.hljs,
.hljs-subst {
  color: #f8f8f2;
}

.hljs-title,
.hljs-attr,
.hljs-meta-keyword {
  font-style: italic;
  color: #50fa7b;
}

.hljs-string,
.hljs-meta,
.hljs-name,
.hljs-type,
.hljs-symbol,
.hljs-bullet,
.hljs-addition,
.hljs-variable,
.hljs-template-tag,
.hljs-template-variable {
  color: #f1fa8c;
}

.hljs-comment,
.hljs-quote,
.hljs-deletion {
  color: #6272a4;
}

.hljs-keyword,
.hljs-selector-tag,
.hljs-literal,
.hljs-title,
.hljs-section,
.hljs-doctag,
.hljs-type,
.hljs-name,
.hljs-strong {
  font-weight: bold;
}

.hljs-literal,
.hljs-number {
  color: #bd93f9;
}

.hljs-emphasis {
  font-style: italic;
}
pre,code {
	-webkit-border-radius: 0.5em;
	border-radius: 0.5em;
}
code.inline {
  -webkit-border-radius: 0.25em;
  border-radius: 0.25em;
  margin-bottom: -0.3em;
  top: 0.1em;
  position: relative;

  display: inline-block;
	padding:0 0.3em;
	font-size:0.9em;
}
blockquote, dl {
	background: #FFF;
  padding: 0.75em 0 0.75em 1em;
  -webkit-border-radius: 0.25em;
  border-radius: 0.25em;
  border-left: 0.3em solid #555;
}

.footnote-definition{
	background: #FFF;
	padding: 0.75em 0 0.75em 1em;
	-webkit-border-radius: 0.25em;
	border-radius: 0.25em;
	border-left: 0.3em solid #555;  
}
		</style>
		<script type="module">
    import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
    mermaid.initialize({ startOnLoad: true });
  </script>
  <script>
  document.addEventListener('DOMContentLoaded', (event) => {
	document.querySelectorAll('code.inline').forEach((el) => {
	  hljs.highlightElement(el);
	});
  });
  
  </script>
  <script type="text/x-mathjax-config">
  MathJax.Hub.Config({
    tex2jax: {
      inlineMath: [ ['$','$'], ["\\(","\\)"] ],
      processEscapes: true
    }
  });
</script>
    
<script type="text/javascript"
        src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
</script>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.2/css/all.min.css" crossorigin="anonymous" referrerpolicy="no-referrer" />
		</body>
	</html>`)

}
