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
		defer func() { // TODO: перенести в отдельную мидлварь.
			p := recover()
			if p != nil {

				c.logger.WarnContext(
					r.Context(), "panic detected",
					slog.Any("panic", p),
				)
			}
		}()

		if !c.debug {
			if next != nil {
				next.ServeHTTP(w, r)
			}

			return
		}

		// Сделано специально для того чтобы получать тут ид трассировки а также иметь информацию о оверхеде с логирования.
		ctx, span := c.tracer.Start(r.Context(), "api server logging")
		defer span.End()

		r = r.WithContext(ctx)

		var (
			responseData = "ignoring"
			requestData  = "ignoring"
		)

		if r.URL.Path != "/api/system/import/archive" { // FIXME: вынести в конфигурацию или проверять по типу контента.
			requestDataRaw, err := io.ReadAll(r.Body)
			if err != nil {
				c.logger.ErrorContext(
					r.Context(), "read request to log",
					slog.Any("error", err),
				)
			}

			requestData = string(requestDataRaw)

			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewReader(requestDataRaw))
		}

		rw := newResponseWrapper(w)

		if next != nil {
			next.ServeHTTP(rw, r)
		}

		if !strings.HasPrefix(r.URL.Path, "/api/file") { // FIXME: вынести в конфигурацию или проверять по типу контента.
			responseData = rw.body.String()
		}

		c.logger.DebugContext(
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
