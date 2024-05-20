package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertItalic(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.ItalicMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "i",
		Contents: c.Convert(elm.GetContent()),
	}
}
