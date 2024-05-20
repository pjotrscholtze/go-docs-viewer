package converter

import (
	"strings"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertEmoji(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.EmojiMarkdownElement)
	if strings.HasPrefix(elm.GetContent(), "fa-") {
		return &htmlwrapper.HTMLElm{
			Tag: "i",
			Attrs: map[string]string{
				"class": "fa-solid " + elm.GetContent(),
			},
		}
	} else {
		return &htmlwrapper.HTMLElm{
			Tag: "span",
			Contents: []htmlwrapper.Elm{
				htmlwrapper.Text(elm.AsMarkdownString()),
			},
		}
	}
}
