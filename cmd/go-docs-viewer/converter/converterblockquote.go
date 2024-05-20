package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertBlockquote(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.BlockQuoteMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag: "blockquote",
		Attrs: map[string]string{
			"class": "blockquote",
		},
		Contents: []htmlwrapper.Elm{
			htmlwrapper.Text(elm.AsMarkdownString()),
		},
	}
}
