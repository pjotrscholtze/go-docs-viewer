package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertCodeblock(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.CodeBlockMarkdownElement)
	mermaidJS := "mermaid"
	if elm.FirstLine()[:min(len(mermaidJS), len(elm.FirstLine()))] == mermaidJS {
		return &htmlwrapper.HTMLElm{
			Tag: "pre",
			Attrs: map[string]string{
				"class": "mermaid",
			},
			Contents: []htmlwrapper.Elm{
				htmlwrapper.Text(elm.WithoutFirstLine()),
			},
		}
	} else {
		return &htmlwrapper.HTMLElm{
			Tag: "pre",
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag: "code",
					Contents: []htmlwrapper.Elm{
						htmlwrapper.Text(elm.WithoutFirstLine()),
					},
				},
			},
		}
	}
}
