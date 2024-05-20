package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertInlineCode(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.InlineCodeMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag: "code",
		Attrs: map[string]string{
			"class": "inline",
		},
		Contents: []htmlwrapper.Elm{
			htmlwrapper.Text(elm.GetContent()),
		},
	}
}
