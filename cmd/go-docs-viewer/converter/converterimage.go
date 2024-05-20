package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertImage(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.ImageMarkdownElement)
	alt := elm.GetAlt()
	attrs := map[string]string{
		"src": elm.GetUrl(),
	}
	if alt != "" {
		attrs["alt"] = alt
	}
	if elm.GetTitle() != nil {
		attrs["title"] = *elm.GetTitle()
	}

	return &htmlwrapper.HTMLElm{
		Tag:      "img",
		Attrs:    attrs,
		Contents: []htmlwrapper.Elm{},
	}
}
