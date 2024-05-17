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
		//
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
	return &htmlwrapper.MultiElm{
		Contents: contents,
	}, outActive
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
			////
			&htmlwrapper.HTMLElm{
				Tag: "div",
				Attrs: map[string]string{
					"id": "content",
				},
				Contents: []htmlwrapper.Elm{
					// bootstrap.NavBar("", bootstrap.BsColorLight, bootstrap.BsLocationNormal, nil, &htmlwrapper.MultiElm{
					// 	Contents: []htmlwrapper.Elm{
					// 		&htmlwrapper.HTMLElm{
					// 			Tag: "button",
					// 			Attrs: map[string]string{
					// 				"type":  "button",
					// 				"id":    "sidebarCollapse",
					// 				"class": "btn btn-info",
					// 			},
					// 			Contents: []htmlwrapper.Elm{
					// 				&htmlwrapper.HTMLElm{
					// 					Tag: "i",
					// 					Attrs: map[string]string{
					// 						"class": "fas fa-align-left",
					// 					},
					// 					Contents: []htmlwrapper.Elm{
					// 						htmlwrapper.Text("Toggle sidebar"),
					// 					},
					// 				},
					// 			},
					// 		},
					// 	},
					// }),
					&htmlwrapper.HTMLElm{
						Tag: "h1",
						Contents: []htmlwrapper.Elm{
							htmlwrapper.Text(title),
						},
					},
					// htmlwrapper.Text("This is an example of a simple admin dashboard"),
					content,
				},
			},
			//

		},
	}
}
