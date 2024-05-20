package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertParagraph(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.ParagraphMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "p",
		Contents: c.Convert(elm.GetContent()),
	}
}
