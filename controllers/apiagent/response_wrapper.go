package apiagent

import (
	"bytes"
	"net/http"
)

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

	return rw.origin.Write(data) //nolint:wrapcheck // не модифицируем
}

func (rw *responseWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.origin.WriteHeader(statusCode)
}
