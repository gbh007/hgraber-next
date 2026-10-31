package agent

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/gbh007/hgraber-next/internal/agent/config"
	hgconfig "github.com/gbh007/hgraber-next/internal/config"

	"go.opentelemetry.io/otel/trace"
)

func parseConfig[T any](defaultParsers func() *T) (config.AgentConfig[T], bool, error) {
	configPath := flag.String("config", "config.yaml", "path to config")
	scan := flag.Bool("scan", false, "scan zip file to register in db")
	useEnv := flag.Bool("use-env", false, "use env config")
	flag.Parse()

	c, err := hgconfig.ImportConfig(*configPath, *useEnv, config.AgentDefaultConfigWrapped(defaultParsers))

	return c, *scan, err
}

func initLogger[T any](cfg config.AgentConfig[T]) *slog.Logger {
	slogOpt := &slog.HandlerOptions{
		AddSource: cfg.Log.IncludeSource,
		Level:     cfg.Log.SlogLevel(),
	}

	return slog.New(
		logHandler{
			Handler: slog.NewJSONHandler(
				os.Stderr,
				slogOpt,
			),
		},
	)
}

// TODO: в случае использования групп реализовать более безопасно.
type logHandler struct {
	slog.Handler
}

func (lh logHandler) Handle(ctx context.Context, r slog.Record) error {
	snapContext := trace.SpanContextFromContext(ctx)
	if snapContext.HasTraceID() {
		r.AddAttrs(slog.String("trace_id", snapContext.TraceID().String()))
	}

	return lh.Handler.Handle(ctx, r)
}

func (lh logHandler) WithGroup(name string) slog.Handler {
	return logHandler{
		Handler: lh.Handler.WithGroup(name),
	}
}

func (lh logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return logHandler{
		Handler: lh.Handler.WithAttrs(attrs),
	}
}
