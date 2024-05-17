package view

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/filetree"
)

func StringToId(in string) string {
	id := in
	for _, old := range []string{"#"} {
		id = strings.ReplaceAll(id, old, "")
	}

	id = strings.Trim(id, " \t")
	for _, old := range []string{" ", "_", ".", "/"} {
		id = strings.ReplaceAll(id, old, "-")
	}

	return strings.ToLower(id)
}
func StringToName(in string) string {
	id := in
	for _, old := range []string{"#"} {
		id = strings.ReplaceAll(id, old, "")
	}

	id = strings.Trim(id, " \t")
	for _, old := range []string{" ", "_", ".", "/"} {
		id = strings.ReplaceAll(id, old, " ")
	}

	if len(id) > 0 {
		id = strings.ToUpper(id[:1]) + id[1:]
	}

	return id
}

func Error(errorMsg string, w http.ResponseWriter) {
	fmt.Fprint(w, `<!doctype html>
	<html lang="en">
	  <head>
	  <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
	  <meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Error while rending the page!</title>
	  </head>
	  <body>
		<h1>Error while rending the page!</h1>
		<pre>`+errorMsg+`</pre>
	  </body>
	</html>`)

}

func NotFound(w http.ResponseWriter) {
	fmt.Fprint(w, `<!doctype html>
	<html lang="en">
	  <head>
	  <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
	  <meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Page not found!</title>
	  </head>
	  <body>
		<h1>Page not found!</h1>
	  </body>
	</html>`)

}

func Menu(currentPath, basePath string, tree []filetree.FileEntry) (htmlwrapper.Elm, bool) {
	contents := []htmlwrapper.Elm{}
	outActive := false
	for _, node := range tree {
		path := "/" + node.Path[len(basePath):]
		id := StringToId(path)
		if node.IsDir {
			items, active := Menu(currentPath, basePath, node.Contents)
			extraCSS, ulCSS := "", "collapse "
			expanded := "false"
			if active {
				outActive = true
				extraCSS = " show"
				ulCSS = "collapse show "
				expanded = "true"
			}

			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "li",
					Attrs: map[string]string{
						"class": "active",
					},
					Contents: []htmlwrapper.Elm{
						&htmlwrapper.HTMLElm{
							Tag: "a",
							Attrs: map[string]string{
								"href":          "#submenu-" + id,
								"data-toggle":   "collapse",
								"aria-expanded": expanded,
								"class":         "dropdown-toggle" + extraCSS,
							},
							Contents: []htmlwrapper.Elm{
								htmlwrapper.Text(StringToName(node.Name)),
							},
						},
						&htmlwrapper.HTMLElm{
							Tag: "ul",
							Attrs: map[string]string{
								"class": ulCSS + "list-unstyled",
								"id":    "submenu-" + id,
							},
							Contents: []htmlwrapper.Elm{
								items,
							},
						},
					},
				})
		} else {
			if path == currentPath {
				outActive = true
			}

			base := filepath.Base(node.Name)
			ext := filepath.Ext(base)
			if ext != "" {
				base = base[:len(base)-len(ext)]
			}

			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "li",
					Contents: []htmlwrapper.Elm{
						&htmlwrapper.HTMLElm{
							Tag: "a",
							Attrs: map[string]string{
								"href": path,
							},
							Contents: []htmlwrapper.Elm{
								htmlwrapper.Text(StringToName(base)),
							}}}})
		}
	}

	return &htmlwrapper.MultiElm{Contents: contents}, outActive
}

func Page(currentPath, basePath string, tree []filetree.FileEntry, content htmlwrapper.Elm) (string, htmlwrapper.Elm) {
	menu, _ := Menu(currentPath, basePath, tree)
	base := filepath.Base(currentPath)
	ext := filepath.Ext(base)
	if ext != "" {
		base = base[:len(base)-len(ext)]
	}

	title := StringToName(base)
	return title, &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "wrapper",
		},
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag: "nav",
				Attrs: map[string]string{
					"id": "sidebar",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "div",
						Attrs: map[string]string{
							"class": "sidebar-header",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.HTMLElm{
								Tag:   "h3",
								Attrs: map[string]string{},
								Contents: []htmlwrapper.Elm{
									htmlwrapper.Text("Admin Panel"),
								},
							},
						},
					},
					&htmlwrapper.HTMLElm{
						Tag: "div",
						Attrs: map[string]string{
							"class": "list-unstyled components",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.HTMLElm{
								Tag: "p",
								Contents: []htmlwrapper.Elm{
									htmlwrapper.Text("Dummy Heading"),
								},
							},
							menu,
						},
					},
				},
			},
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"id": "content",
				},
				Contents: []htmlwrapper.Elm{
					&htmlwrapper.HTMLElm{
						Tag: "h1",
						Contents: []htmlwrapper.Elm{
							htmlwrapper.Text(title),
						},
					},
					content,
				},
			},
		},
	}
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
