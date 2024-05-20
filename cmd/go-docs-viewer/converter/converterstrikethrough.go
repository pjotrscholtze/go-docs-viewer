package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertStrikethrough(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.StrikethroughMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "s",
		Contents: c.Convert(elm.GetContent()),
	}
}
