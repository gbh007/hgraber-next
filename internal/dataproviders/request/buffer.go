package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// requestBuffer запрашивает данные по урле и возвращает их в виде буффера
func (r *Requester) requestBuffer(ctx context.Context, URL string, headers http.Header, body io.Reader) (*bytes.Buffer, string, error) {
	buff := new(bytes.Buffer)

	var (
		req *http.Request
		err error
	)

	if body != nil {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, URL, body)
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, URL, http.NoBody)
	}

	if err != nil {
		return nil, "", err
	}

	if len(headers) > 0 {
		for key, values := range headers {
			for _, v := range values {
				req.Header.Add(key, v)
			}
		}
	}

	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", r.defaultUserAgent)
	}

	// выполняем запрос
	response, err := r.client.Do(req)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		closeErr := response.Body.Close()
		if closeErr != nil {
			r.logger.ErrorContext(ctx, closeErr.Error())
		}
	}()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, "", fmt.Errorf("load %s: unsuccess status: %s", URL, response.Status)
	}

	_, err = buff.ReadFrom(response.Body)
	if err != nil {
		return nil, "", err
	}

	if response.Request != nil &&
		response.Request.URL != nil &&
		response.Request.URL.String() != URL {
		URL = response.Request.URL.String()
	}

	return buff, URL, nil
}
