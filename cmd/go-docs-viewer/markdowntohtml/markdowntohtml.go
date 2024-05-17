package markdowntohtml

import (
	"fmt"
	"strings"

	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-docs-viewer/cmd/go-docs-viewer/view"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func HeadingToId(heading entity.HeaderMarkdownElement) string {
	return view.StringToId(heading.AsMarkdownString())
}

func convertListTuple(listTuple []entity.ListTuple) []htmlwrapper.Elm {
	res := []htmlwrapper.Elm{}
	for _, entry := range listTuple {
		if entry.List != nil {
			res = append(res, Convert([]entity.MarkdownElement{entry.List})...)
		} else {
			res = append(res, Convert([]entity.MarkdownElement{entry.ListItem})...)
		}
	}
	return res
}

func Convert(doc []entity.MarkdownElement) []htmlwrapper.Elm {
	contents := []htmlwrapper.Elm{}
	for _, element := range doc {
		switch element.Kind() {
		case entity.ElementKindText:
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "span",
					Contents: []htmlwrapper.Elm{htmlwrapper.Text(element.AsMarkdownString())},
				},
			)
			break
		case entity.ElementKindHeader:
			heading := element.(entity.HeaderMarkdownElement)
			id := HeadingToId(heading)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: fmt.Sprintf("h%d", heading.GetHeadingLevel()),
					Attrs: map[string]string{
						"id": id,
					},
					Contents: Convert(heading.GetChildren()),
				},
			)
			break
		case entity.ElementKindHorizontalLine:
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "hr",
				},
			)
			break
		case entity.ElementKindBold:
			elm := element.(entity.BoldMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "strong",
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindItalic:
			elm := element.(entity.ItalicMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "i",
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindStrikethrough:
			elm := element.(entity.StrikethroughMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "s",
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindLink:
			elm := element.(entity.LinkMarkdownElement)
			title := elm.GetTitle()
			attrs := map[string]string{
				"href": elm.GetUrl(),
			}
			if title != nil {
				attrs["title"] = *title
			}
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "a",
					Attrs:    attrs,
					Contents: []htmlwrapper.Elm{htmlwrapper.Text(elm.GetContent())},
				},
			)
			break
		case entity.ElementKindCodeblock:
			elm := element.(entity.CodeBlockMarkdownElement)
			mermaidJS := "mermaid"
			if elm.FirstLine()[:min(len(mermaidJS), len(elm.FirstLine()))] == mermaidJS {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "pre",
						Attrs: map[string]string{
							"class": "mermaid",
						},
						Contents: []htmlwrapper.Elm{
							htmlwrapper.Text(elm.WithoutFirstLine()),
						},
					},
				)
			} else {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "pre",
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.HTMLElm{
								Tag: "code",
								Contents: []htmlwrapper.Elm{
									htmlwrapper.Text(elm.WithoutFirstLine()),
								},
							},
						},
					},
				)
			}
			break
		case entity.ElementKindListItem:
			elm := element.(entity.ListItemElementMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "li",
					Contents: Convert(elm.GetContentMarkdownElements()),
				},
			)
			break
		case entity.ElementKindList:
			elm := element.(entity.ListElementMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "ul",
					Contents: convertListTuple(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindHighlight:
			elm := element.(entity.HighlightMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "mark",
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindEmoji:
			elm := element.(entity.EmojiMarkdownElement)
			if strings.HasPrefix(elm.GetContent(), "fa-") {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "i",
						Attrs: map[string]string{
							"class": "fa-solid " + elm.GetContent(),
						},
					},
				)
			} else {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "span",
						Contents: []htmlwrapper.Elm{
							htmlwrapper.Text(elm.AsMarkdownString()),
						},
					},
				)
			}
			break
		case entity.ElementKindTable:
			elm := element.(entity.TableElementMarkdownElement)
			header := elm.Header()
			heading := []htmlwrapper.Elm{}
			body := []htmlwrapper.Elm{}
			if header != nil {
				// heading
				// headingCells := []htmlwrapper.Elm{}
				for _, cell := range header.Cells {
					heading = append(heading, bootstrap.TableCell(false, bootstrap.BsTableCellKindNormal, 1, bootstrap.BsTableColorDark,
						Convert([]entity.MarkdownElement{cell}),
					))
				}
				// heading = append(heading, headingCells)
			}
			for _, row := range elm.Rows() {
				rowCells := []htmlwrapper.Elm{}
				for _, cell := range row.Cells {
					rowCells = append(rowCells, bootstrap.TableCell(false, bootstrap.BsTableCellKindNormal, 1, bootstrap.BsTableColorDefault,
						// htmlwrapper.Text(cell),
						Convert([]entity.MarkdownElement{cell}),
					))
				}
				// &htmlwrapper.MultiElm{rowCells}
				body = append(body, bootstrap.TableRow(false, bootstrap.BsTableColorDefault, rowCells))
			}
			contents = append(contents, bootstrap.Table(true, false, bootstrap.BsTableColorDefault, bootstrap.BsTableBorderColorDefault, bootstrap.BsTablSizeLarge, heading, body, nil, nil))
			// contents = append(contents,
			// 	&htmlwrapper.HTMLElm{
			// 		Tag:      "mark",
			// 		Contents: Convert(elm.GetContent()),
			// 	},
			// )
			break
		case entity.ElementKindInlineCode:
			elm := element.(entity.InlineCodeMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					// Tag: "pre",
					// Attrs: map[string]string{
					// 	"class": "inline",
					// },
					// Contents: []htmlwrapper.Elm{
					// 	&htmlwrapper.HTMLElm{
					Tag: "code",
					Attrs: map[string]string{
						"class": "inline",
					},
					Contents: []htmlwrapper.Elm{
						htmlwrapper.Text(elm.GetContent()),
						// 	},
						// },
					},
				},
			)
			// contents = append(contents,
			// 		&htmlwrapper.HTMLElm{
			// 			Tag: "code",
			// 			Contents: []htmlwrapper.Elm{
			// 				htmlwrapper.Text(elm.AsMarkdownString()),
			// 			},
			// 		},
			// 	)
			break
		case entity.ElementKindBlockquote:
			elm := element.(entity.BlockQuoteMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "blockquote",
					Attrs: map[string]string{
						"class": "blockquote",
					},
					Contents: []htmlwrapper.Elm{
						htmlwrapper.Text(elm.AsMarkdownString()),
					},
				},
			)
			break
		case entity.ElementKindImage:
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

			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "img",
					Attrs:    attrs,
					Contents: []htmlwrapper.Elm{},
				},
			)
			break
		case entity.ElementKindCheckbox:
			elm := element.(entity.CheckboxMarkdownElement)
			attrs := map[string]string{
				"type":     "checkbox",
				"disabled": "disabled",
			}
			if elm.GetCheckContent() != " " {
				attrs["checked"] = "checked"
			}
			contents = append(contents,
				&htmlwrapper.MultiElm{
					Contents: []htmlwrapper.Elm{
						&htmlwrapper.HTMLElm{
							Tag:   "input",
							Attrs: attrs,
						},
					},
				},
			)
			break
		case entity.ElementKindTermDefinitionLine:
			elm := element.(entity.TermDefinitionElementMarkdownElement)
			content := Convert(elm.GetContent())

			if len(content) > 1 {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "dl",
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.HTMLElm{
								Tag: "dt",
								Attrs: map[string]string{
									"class": "term-definition",
								},
								Contents: content[:1],
							},
							&htmlwrapper.HTMLElm{
								Tag: "dd",
								Attrs: map[string]string{
									"class": "term-definition",
								},
								Contents: content[1:],
							},
						},
					},
				)
			} else {
				contents = append(contents,
					&htmlwrapper.HTMLElm{
						Tag: "dl",
						Attrs: map[string]string{
							"class": "term-definition",
						},
						Contents: []htmlwrapper.Elm{
							&htmlwrapper.HTMLElm{
								Tag:      "dd",
								Contents: content,
							},
						},
					},
				)
			}
			break
		case entity.ElementKindParagraph:
			elm := element.(entity.ParagraphMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "p",
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindFootnote:
			elm := element.(entity.FootnoteMarkdownElement)
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "sup",
					Attrs: map[string]string{
						"class": "footnote",
					},
					Contents: Convert(elm.GetContent()),
				},
			)
			break
		case entity.ElementKindFootnoteDefinition:
			elm := element.(entity.FootnoteDefinitionMarkdownElement)
			c := []htmlwrapper.Elm{
				&htmlwrapper.HTMLElm{
					Tag: "sup",
					Attrs: map[string]string{
						"class": "footnote",
					},
					Contents: []htmlwrapper.Elm{
						&htmlwrapper.MultiElm{
							Contents: Convert(elm.GetContent()),
						},
					},
				},
				htmlwrapper.Text(": "),
			}
			c = append(c, Convert(elm.GetDefinition())...)

			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag: "div",
					Attrs: map[string]string{
						"class": "footnote-definition",
					},
					Contents: c,
				},
			)
			break
		case entity.ElementKindGroup:
			elm := element.(*entity.GroupElement)
			contents = append(contents,
				Convert(elm.Contents)...,
			)
			break
		default:
			contents = append(contents,
				&htmlwrapper.HTMLElm{
					Tag:      "pre",
					Contents: []htmlwrapper.Elm{htmlwrapper.Text(element.AsMarkdownString())},
				},
			)
		}
	}
	return contents
}
