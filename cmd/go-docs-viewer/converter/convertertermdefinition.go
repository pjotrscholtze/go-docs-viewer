package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertTermdefinition(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.TermDefinitionElementMarkdownElement)
	content := c.Convert(elm.GetContent())

	if len(content) > 1 {
		return &htmlwrapper.HTMLElm{
			Tag: "dl",
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag: "dt",
					Attrs: map[string]string{
						"class": "term-definition",
					},
					Contents: content[:1],
				},
				&htmlwrapper.HTMLElm{
					Tag: "dd",
					Attrs: map[string]string{
						"class": "term-definition",
					},
					Contents: content[1:],
				},
			},
		}
	} else {
		return &htmlwrapper.HTMLElm{
			Tag: "dl",
			Attrs: map[string]string{
				"class": "term-definition",
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag:      "dd",
					Contents: content,
				},
			},
		}
	}
}
