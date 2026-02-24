package components

import (
	"persephone/internal/ui/theme"

	"github.com/rivo/tview"
)

func NewTableList() *tview.List {

	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle(" 📋 Tables ")
	list.SetBorderColor(theme.PrimaryColor)
	list.SetTitleColor(theme.PrimaryColor)
	list.ShowSecondaryText(false)
	list.SetHighlightFullLine(true)
	list.SetBorderPadding(0, 0, 1, 1)
	list.SetSelectedBackgroundColor(theme.AccentColor)

	return list
}
