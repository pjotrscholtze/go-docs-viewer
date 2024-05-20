package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertText(element entity.MarkdownElement) htmlwrapper.Elm {
	return &htmlwrapper.HTMLElm{
		Tag:      "span",
		Contents: []htmlwrapper.Elm{htmlwrapper.Text(element.AsMarkdownString())},
	}
}
