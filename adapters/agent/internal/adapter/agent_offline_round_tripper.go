package adapter

import (
	"errors"
	"fmt"
	"net/http"
	"syscall"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
)

type agentOfflineRoundTripper struct {
	next http.RoundTripper
}

func (rt agentOfflineRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := rt.next.RoundTrip(req)

	if errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EHOSTDOWN) ||
		errors.Is(err, syscall.EHOSTUNREACH) {
		err = fmt.Errorf("%w: %w", agentmodel.ErrAgentAPIOffline, err)
	}

	return res, err
}
