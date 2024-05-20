package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertHighlight(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.HighlightMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "mark",
		Contents: c.Convert(elm.GetContent()),
	}
}
