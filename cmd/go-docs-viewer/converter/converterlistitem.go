package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertListItem(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.ListItemElementMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "li",
		Contents: c.Convert(elm.GetContentMarkdownElements()),
	}
}
