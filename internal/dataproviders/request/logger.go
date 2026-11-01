package request

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type requestErrorLogger struct {
	logger *slog.Logger
	next   http.RoundTripper
}

func (h *requestErrorLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())

	var (
		responseBody []byte
		err          error
		args         []any
	)

	resp, err := h.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 400 {
		return resp, nil
	}

	args = append(
		args,
		slog.String("status", strconv.Itoa(resp.StatusCode)),
	)

	if len(resp.Header) > 0 {
		headers := make(map[string]string)

		for k, v := range resp.Header {
			headers[k] = strings.Join(v, ";")
		}

		args = append(
			args,
			slog.Any("headers", headers),
		)
	}

	if resp.Body != nil &&
		strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "application/json") ||
		strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "text/") {
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		resp.Body = io.NopCloser(bytes.NewReader(responseBody))
	}

	if len(responseBody) > 0 {
		args = append(
			args,
			slog.String("body", string(responseBody)),
		)
	}

	h.logger.ErrorContext(
		req.Context(), "error on request request",
		args...,
	)

	return resp, nil
}
