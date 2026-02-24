package ui

import (
	"database/sql"
	"persephone/internal/repository"
	"persephone/internal/service"
	"persephone/internal/ui/page"

	"github.com/rivo/tview"
)

type App struct {
	app  *tview.Application
	conn *sql.DB
}

func NewApp(conn *sql.DB) *App {
	return &App{
		app:  tview.NewApplication(),
		conn: conn,
	}
}

func (a *App) Run() error {

	tableRepository := repository.NewTableRepository(a.conn)
	columnRepository := repository.NewColumnRepository(a.conn)
	indexRepository := repository.NewIndexRepository(a.conn)

	tableService := service.NewTableService(
		tableRepository,
		columnRepository,
		indexRepository,
	)

	mainPage := page.NewMainPage(
		tableService,
		a.app,
	)

	p, err := mainPage.Build()
	if err != nil {
		return err
	}

	return a.app.SetRoot(p, true).
		EnableMouse(true).
		EnablePaste(true).
		Run()
}

func (a *App) Stop() {
	a.app.Stop()
}
