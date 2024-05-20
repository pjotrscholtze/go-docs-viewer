package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertGroup(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(*entity.GroupElement)
	return &htmlwrapper.MultiElm{
		Contents: c.Convert(elm.Contents),
	}
}
