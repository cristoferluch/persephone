package components

import (
	"persephone/internal/ui/theme"

	"github.com/rivo/tview"
)

func NewIndexTable() *tview.Table {

	table := tview.NewTable()
	table.SetBorders(true)
	table.SetBorder(true)
	table.SetTitle(" 🔑 Indexes ")
	table.SetBorderColor(theme.BorderColor)
	table.SetTitleColor(theme.BorderColor)
	table.SetFixed(1, 0)

	return table
}
