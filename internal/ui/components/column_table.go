package components

import (
	"persephone/internal/ui/theme"

	"github.com/rivo/tview"
)

func NewColumnTable() *tview.Table {

	table := tview.NewTable()
	table.SetBorders(true)
	table.SetBorder(true)
	table.SetTitle(" 📑 Columns ")
	table.SetBorderColor(theme.HeaderColor)
	table.SetTitleColor(theme.HeaderColor)
	table.SetFixed(1, 0)

	return table
}
