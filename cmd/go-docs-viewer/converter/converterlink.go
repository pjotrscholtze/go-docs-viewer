package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertLink(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.LinkMarkdownElement)
	title := elm.GetTitle()
	attrs := map[string]string{
		"href": elm.GetUrl(),
	}
	if title != nil {
		attrs["title"] = *title
	}
	return &htmlwrapper.HTMLElm{
		Tag:      "a",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{htmlwrapper.Text(elm.GetContent())},
	}
}
