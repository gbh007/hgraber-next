package apiserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	"hgnext/internal/controllers/apiserver/internal/static"
)

func (c *Controller) Name() string {
	return "api server"
}

func (c *Controller) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	mux := http.NewServeMux()

	// обработчик статики
	if c.staticDir != "" {
		mux.Handle("/", http.FileServer(http.Dir(c.staticDir)))
	} else {
		mux.Handle("/", http.FileServer(http.FS(static.StaticDir)))
	}

	mux.Handle("/api/", c.logIO(cors(c.ogenServer)))

	server := &http.Server{
		Handler: mux,
		Addr:    c.serverAddr,
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
