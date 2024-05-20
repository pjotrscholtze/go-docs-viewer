package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertCheckbox(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.CheckboxMarkdownElement)
	attrs := map[string]string{
		"type":     "checkbox",
		"disabled": "disabled",
	}
	if elm.GetCheckContent() != " " {
		attrs["checked"] = "checked"
	}
	return &htmlwrapper.MultiElm{
		Contents: []htmlwrapper.Elm{
			&htmlwrapper.HTMLElm{
				Tag:   "input",
				Attrs: attrs,
			},
		},
	}
}
