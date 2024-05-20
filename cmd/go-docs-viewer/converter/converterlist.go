package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertListTuple(listTuple []entity.ListTuple) []htmlwrapper.Elm {
	res := []htmlwrapper.Elm{}
	for _, entry := range listTuple {
		if entry.List != nil {
			res = append(res, c.Convert([]entity.MarkdownElement{entry.List})...)
		} else {
			res = append(res, c.Convert([]entity.MarkdownElement{entry.ListItem})...)
		}
	}
	return res
}

func (c *converter) convertList(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.ListElementMarkdownElement)
	return &htmlwrapper.HTMLElm{
		Tag:      "ul",
		Contents: c.convertListTuple(elm.GetContent()),
	}
}
