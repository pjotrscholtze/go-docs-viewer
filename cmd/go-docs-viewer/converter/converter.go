package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/view"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

type converter struct {
	converter map[string]func(element entity.MarkdownElement) htmlwrapper.Elm
}
type Converter interface {
	Convert(doc []entity.MarkdownElement) []htmlwrapper.Elm
}

func (c *converter) Convert(doc []entity.MarkdownElement) []htmlwrapper.Elm {
	contents := []htmlwrapper.Elm{}
	for _, element := range doc {
		fn, ok := c.converter[element.Kind()]
		if !ok {
			// default fn here
			continue
		}
		contents = append(contents, fn(element))

	}

	return contents
}

func (c *converter) HeadingToId(heading entity.HeaderMarkdownElement) string {
	return view.StringToId(heading.AsMarkdownString())
}

func NewConverter() Converter {
	c := &converter{}
	c.converter = map[string]func(element entity.MarkdownElement) htmlwrapper.Elm{
		entity.ElementKindText:               c.convertText,
		entity.ElementKindHeader:             c.convertHeader,
		entity.ElementKindHorizontalLine:     c.convertHorizontalLine,
		entity.ElementKindBold:               c.convertBold,
		entity.ElementKindItalic:             c.convertItalic,
		entity.ElementKindLink:               c.convertLink,
		entity.ElementKindStrikethrough:      c.convertStrikethrough,
		entity.ElementKindCodeblock:          c.convertCodeblock,
		entity.ElementKindListItem:           c.convertListItem,
		entity.ElementKindList:               c.convertList,
		entity.ElementKindHighlight:          c.convertHighlight,
		entity.ElementKindEmoji:              c.convertEmoji,
		entity.ElementKindTable:              c.convertTable,
		entity.ElementKindInlineCode:         c.convertInlineCode,
		entity.ElementKindBlockquote:         c.convertBlockquote,
		entity.ElementKindImage:              c.convertImage,
		entity.ElementKindCheckbox:           c.convertCheckbox,
		entity.ElementKindTermDefinitionLine: c.convertTermdefinition,
		entity.ElementKindParagraph:          c.convertParagraph,
		entity.ElementKindFootnote:           c.convertFootnote,
		entity.ElementKindFootnoteDefinition: c.convertFootnoteDefinition,
		entity.ElementKindGroup:              c.convertGroup,
	}
	return c
}
