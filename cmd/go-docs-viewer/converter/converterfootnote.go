package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertFootnote(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.FootnoteMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag: "sup",
		Attrs: map[string]string{
			"class": "footnote",
		},
		Contents: c.Convert(elm.GetContent()),
	}
}
