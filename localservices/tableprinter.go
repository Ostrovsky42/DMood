package localservices

import (
	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
	"os"
)


func NewPrinter()*tableprinter.Printer  {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	return printer
}