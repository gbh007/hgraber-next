package apiserver

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func (c *Controller) logIO(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !c.debug {
			if next != nil {
				next.ServeHTTP(w, r)
			}

			return
		}

		requestData, err := io.ReadAll(r.Body)
		if err != nil {
			c.logger.ErrorContext(
				r.Context(), "read request to log",
				slog.Any("error", err),
			)
		}

		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(requestData))

		rw := newResponseWrapper(w)

		if next != nil {
			next.ServeHTTP(rw, r)
		}

		var responseData = "ignoring"

		// FIXME: не учтено экранирование пароля пользователя и т.п.
		if !strings.HasPrefix(r.URL.Path, "/api/file") {
			responseData = rw.body.String()
		}

		c.logger.DebugContext(
			r.Context(), "http request",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
			slog.Group(
				"request",
				slog.Any("headers", r.Header),
				slog.String("body", string(requestData)),
			),
			slog.Group(
				"response",
				slog.Int("code", rw.statusCode),
				slog.Any("headers", rw.origin.Header()),
				slog.String("body", responseData),
			),
		)
	})
}

type responseWrapper struct {
	origin http.ResponseWriter

	statusCode int
	body       *bytes.Buffer
}

func newResponseWrapper(origin http.ResponseWriter) *responseWrapper {
	return &responseWrapper{
		origin: origin,
		body:   &bytes.Buffer{},
	}
}

func (rw *responseWrapper) Header() http.Header {
	return rw.origin.Header()
}

func (rw *responseWrapper) Write(data []byte) (int, error) {
	_, _ = rw.body.Write(data)

	return rw.origin.Write(data)
}

func (rw *responseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.origin.WriteHeader(statusCode)
}
