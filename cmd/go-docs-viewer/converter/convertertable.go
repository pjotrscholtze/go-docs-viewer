package converter

import (
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/bootstrap"
	"github.com/pjotrscholtze/go-bootstrap/cmd/go-bootstrap/htmlwrapper"
	"github.com/pjotrscholtze/go-markdown/cmd/go-markdown/entity"
)

func (c *converter) convertTable(element entity.MarkdownElement) htmlwrapper.Elm {
	elm := element.(entity.TableElementMarkdownElement)
	header := elm.Header()
	heading := []htmlwrapper.Elm{}
	body := []htmlwrapper.Elm{}
	if header != nil {
		for _, cell := range header.Cells {
			heading = append(heading, bootstrap.TableCell(false, bootstrap.BsTableCellKindNormal, 1, bootstrap.BsTableColorDark,
				c.Convert([]entity.MarkdownElement{cell}),
			))
		}
	}
	for _, row := range elm.Rows() {
		rowCells := []htmlwrapper.Elm{}
		for _, cell := range row.Cells {
			rowCells = append(rowCells, bootstrap.TableCell(false, bootstrap.BsTableCellKindNormal, 1, bootstrap.BsTableColorDefault,
				c.Convert([]entity.MarkdownElement{cell}),
			))
		}
		body = append(body, bootstrap.TableRow(false, bootstrap.BsTableColorDefault, rowCells))
	}
	return bootstrap.Table(true, false, bootstrap.BsTableColorDefault, bootstrap.BsTableBorderColorDefault, bootstrap.BsTablSizeLarge, heading, body, nil, nil)
}
