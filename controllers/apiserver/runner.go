package apiserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (c *Controller) Name() string {
	return "api server"
}

func (c *Controller) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	mux := http.NewServeMux()

	if c.staticDir != "" {
		mux.Handle("/", http.FileServer(http.Dir(c.staticDir)))
	}

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/api/", otelPropagation(c.logIO(cors(c.ogenServer))))

	server := &http.Server{
		Handler:  mux,
		Addr:     c.serverAddr,
		ErrorLog: slog.NewLogLogger(c.logger.Handler(), slog.LevelError),
	}

	go func() {
		defer close(done)

		c.logger.InfoContext(parentCtx, "api server start")
		defer c.logger.InfoContext(parentCtx, "api server stop")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	go func() {
		<-parentCtx.Done()
		c.logger.InfoContext(parentCtx, "stopping api server")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(parentCtx), time.Second*10)
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
