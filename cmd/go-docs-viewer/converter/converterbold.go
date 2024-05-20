package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertBold(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.BoldMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "strong",
		Contents: c.Convert(elm.GetContent()),
	}
}
