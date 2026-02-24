package page

import (
	"persephone/internal/entity"
	"strconv"
	"strings"

	"persephone/internal/service"
	"persephone/internal/ui/components"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Page struct {
	tableService *service.TableService

	app *tview.Application

	pages       *tview.Pages
	searchInput *tview.InputField
	tableList   *tview.List
	columnTable *tview.Table
	indexTable  *tview.Table
}

func NewMainPage(
	tableService *service.TableService,
	app *tview.Application,
) *Page {
	return &Page{
		tableService: tableService,
		pages:        tview.NewPages(),
		app:          app,
	}
}

func (p *Page) Build() (*tview.Pages, error) {

	p.searchInput = components.NewSearchInput()
	p.tableList = components.NewTableList()
	p.columnTable = components.NewColumnTable()
	p.indexTable = components.NewIndexTable()

	if err := p.populateTables(""); err != nil {
		return nil, err
	}

	if err := p.setInputEvents(); err != nil {
		return nil, err
	}

	if err := p.setListEvents(); err != nil {
		return nil, err
	}

	leftPanel := tview.NewFlex().
		AddItem(p.searchInput, 3, 0, true).
		AddItem(p.tableList, 0, 1, true).
		SetDirection(tview.FlexRow)

	rightPanel := tview.NewFlex().
		AddItem(p.indexTable, 0, 1, false).
		AddItem(p.columnTable, 0, 1, false).
		SetDirection(tview.FlexRow)

	layout := tview.NewFlex().
		AddItem(leftPanel, 0, 1, false).
		AddItem(rightPanel, 0, 2, false).
		SetDirection(tview.FlexColumn)

	p.pages.AddPage("form", layout, true, true)

	p.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTAB {
			switch {
			case p.searchInput.HasFocus():
				p.app.SetFocus(p.tableList)
				p.tableList.SetCurrentItem(-1)
			}
		}
		if event.Key() == tcell.KeyCtrlK {
			p.app.SetFocus(p.searchInput)
		}

		return event
	})

	return p.pages, nil
}

func (p *Page) populateTables(search string) error {

	p.tableList.Clear()

	search = strings.TrimSpace(search)

	tables, err := p.tableService.FindAll(search)
	if err != nil {
		return err
	}

	tablesFilter := make([]entity.Table, 0)
	for _, table := range tables {
		if strings.Contains(strings.ToLower(table.Name), strings.ToLower(search)) {
			tablesFilter = append(tablesFilter, table)
		}
	}

	for _, table := range tablesFilter {
		p.tableList.AddItem(" • "+table.Name, table.Name, 0, nil)
	}

	return nil
}

func (p *Page) setInputEvents() error {
	p.searchInput.SetChangedFunc(func(text string) {

		err := p.populateTables(text)
		if err != nil {
			return
		}
	})
	return nil
}

func (p *Page) setListEvents() error {
	p.tableList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {

		if err := p.populateColumnTable(secondaryText); err != nil {
			return
		}
		if err := p.populateIndexTable(secondaryText); err != nil {
			return
		}

	})
	return nil
}

func (p *Page) populateColumnTable(tableName string) error {

	tableName = strings.TrimSpace(tableName)

	p.columnTable.Clear()
	row := 0

	p.columnTable.SetCell(row, 0, tview.NewTableCell("Name").SetExpansion(1))
	p.columnTable.SetCell(row, 1, tview.NewTableCell("Type").SetExpansion(1))
	p.columnTable.SetCell(row, 2, tview.NewTableCell("Length").SetExpansion(1))
	p.columnTable.SetCell(row, 3, tview.NewTableCell("Precision").SetExpansion(1))
	p.columnTable.SetCell(row, 4, tview.NewTableCell("Nullable").SetExpansion(1))
	p.columnTable.SetCell(row, 5, tview.NewTableCell("Primary key").SetExpansion(1))

	columnList, err := p.tableService.GetColumnByTable(tableName)
	if err != nil {
		return err
	}

	for _, column := range columnList {
		row++
		p.columnTable.SetCell(row, 0, tview.NewTableCell(column.Name).SetExpansion(1))
		p.columnTable.SetCell(row, 1, tview.NewTableCell(column.Type).SetExpansion(1))
		p.columnTable.SetCell(row, 2, tview.NewTableCell(strconv.Itoa(column.Length)).SetExpansion(1))
		p.columnTable.SetCell(row, 3, tview.NewTableCell(strconv.Itoa(column.Precision)).SetExpansion(1))
		p.columnTable.SetCell(row, 4, tview.NewTableCell(strconv.FormatBool(column.IsNullable)).SetExpansion(1))
		p.columnTable.SetCell(row, 5, tview.NewTableCell(strconv.FormatBool(column.HasPrimaryKey)).SetExpansion(1))
	}
	p.columnTable.ScrollToBeginning()

	return nil
}

func (p *Page) populateIndexTable(tableName string) error {

	tableName = strings.TrimSpace(tableName)

	p.indexTable.Clear()

	row := 0

	p.indexTable.SetCell(row, 0, tview.NewTableCell("Name").SetExpansion(1))
	p.indexTable.SetCell(row, 1, tview.NewTableCell("Description").SetExpansion(1))
	p.indexTable.SetCell(row, 2, tview.NewTableCell("Keys").SetExpansion(1))

	indexList, err := p.tableService.GetIndexByTable(tableName)
	if err != nil {
		return err
	}

	for _, index := range indexList {
		row++
		p.indexTable.SetCell(row, 0, tview.NewTableCell(index.Name).SetExpansion(1))
		p.indexTable.SetCell(row, 1, tview.NewTableCell(index.Description).SetExpansion(1))
		p.indexTable.SetCell(row, 2, tview.NewTableCell(index.Keys).SetExpansion(1))
	}

	return nil
}
