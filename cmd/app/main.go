package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"
	"persephone/internal/config"
	"persephone/internal/database"
	"persephone/internal/ui"
	"syscall"
)

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	conn, err := database.NewPostgresSQL(*cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	app := ui.NewApp(conn)

	go func() {
		<-ctx.Done()
		app.Stop()
	}()

	if err = app.Run(); err != nil {
		panic(err)
	}

}
