package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gbh007/hgraber-next/application/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	a := server.New()

	err := a.Init(ctx)
	if err != nil {
		a.Logger.ErrorContext(
			ctx, "init application error",
			slog.Any("error", err),
		)

		cancel()

		os.Exit(1) //nolint:gocritic // вызывается вручную до
	}

	err = a.Serve(ctx)
	if err != nil {
		a.Logger.ErrorContext(
			ctx, "serve application error",
			slog.Any("error", err),
		)

		cancel()

		os.Exit(1)
	}
}
