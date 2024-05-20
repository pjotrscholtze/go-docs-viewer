package converter

import (
	"fmt"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertHeader(element entity.MarkdownElement) htmlwrapper.Elm {
	heading := element.(entity.HeaderMarkdownElement)
	id := c.HeadingToId(heading)
	return &htmlwrapper.HTMLElm{
		Tag: fmt.Sprintf("h%d", heading.GetHeadingLevel()),
		Attrs: map[string]string{
			"id": id,
		},
		Contents: c.Convert(heading.GetChildren()),
	}
}
