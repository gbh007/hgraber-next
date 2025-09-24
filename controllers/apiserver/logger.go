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

		// Сделано специально для того чтобы получать тут ид
		// трассировки а также иметь информацию о оверхеде с
		// логирования.
		ctx, span := c.tracer.Start(r.Context(), "api server logging")
		defer span.End()

		r = r.WithContext(ctx)

		var (
			responseData = "ignoring"
			requestData  = "ignoring"
		)

		if strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "application/json") ||
			strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "text/") {
			requestDataRaw, err := io.ReadAll(r.Body)
			if err != nil {
				c.logger.ErrorContext( //nolint:contextcheck // ложно-положительное срабатывание
					r.Context(), "read request to log",
					slog.Any("error", err),
				)
			}

			requestData = string(requestDataRaw)

			err = r.Body.Close()
			if err != nil {
				c.logger.ErrorContext( //nolint:contextcheck // ложное срабатывание
					r.Context(), "close request body",
					slog.Any("error", err),
				)
			}

			r.Body = io.NopCloser(bytes.NewReader(requestDataRaw))
		}

		rw := newResponseWrapper(w)

		if next != nil {
			next.ServeHTTP(rw, r)
		}

		if strings.Contains(strings.ToLower(rw.Header().Get("Content-Type")), "application/json") ||
			strings.Contains(strings.ToLower(rw.Header().Get("Content-Type")), "text/") {
			responseData = rw.body.String()
		}

		c.logger.DebugContext( //nolint:contextcheck // ложно-положительное срабатывание
			r.Context(), "http request",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
			slog.Group(
				"request",
				slog.Any("headers", r.Header),
				slog.String("body", requestData),
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
