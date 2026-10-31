package apiagent

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (c *Controller) Name() string {
	return "api server"
}

func (c *Controller) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	server := &http.Server{ //nolint:gosec // будет исправлено позднее
		Handler:  otelPropagation(c.logIO(cors(c.ogenServer))),
		Addr:     c.addr,
		ErrorLog: slog.NewLogLogger(c.logger.Handler(), slog.LevelError),
	}

	go func() {
		defer close(done)

		c.logger.InfoContext(parentCtx, "api agent server start")
		defer c.logger.InfoContext(parentCtx, "api agent server stop")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	go func() {
		<-parentCtx.Done()
		c.logger.InfoContext(parentCtx, "stopping api agent server")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(parentCtx), time.Second*10) //nolint:mnd,lll,golines // будет исправлено позднее
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	return done, nil
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

func otelPropagation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(
			r.Context(),
			propagation.HeaderCarrier(r.Header),
		)

		if next != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
