package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertFootnoteDefinition(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.FootnoteDefinitionMarkdownElement)
	content := []htmlwrapper.Elm{
		&htmlwrapper.HTMLElm{
			Tag: "sup",
			Attrs: map[string]string{
				"class": "footnote",
			},
			Contents: []htmlwrapper.Elm{
				&htmlwrapper.MultiElm{
					Contents: c.Convert(elm.GetContent()),
				},
			},
		},
		htmlwrapper.Text(": "),
	}
	content = append(content, c.Convert(elm.GetDefinition())...)

	return &htmlwrapper.HTMLElm{
		Tag: "div",
		Attrs: map[string]string{
			"class": "footnote-definition",
		},
		Contents: content,
	}
}
