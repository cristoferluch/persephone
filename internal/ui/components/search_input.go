package components

import (
	"persephone/internal/ui/theme"

	"github.com/rivo/tview"
)

func NewSearchInput() *tview.InputField {

	input := tview.NewInputField()
	input.SetLabel(" 🔍 Search: ")
	input.SetFieldBackgroundColor(theme.BgColor)
	input.SetBorder(true)
	input.SetBorderColor(theme.AccentColor)
	input.SetBorderPadding(0, 0, 1, 1)

	return input
}
