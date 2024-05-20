package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertHorizontalLine(element entity.MarkdownElement) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag: "hr",
	}
}
