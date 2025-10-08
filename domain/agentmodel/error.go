package agentmodel

import "errors"

var (
	ErrAgentAPIOffline         = errors.New("agent: offline")
	ErrAgentAPIUnauthorized    = errors.New("agent: unauthorized")
	ErrAgentAPIForbidden       = errors.New("agent: forbidden")
	ErrAgentAPIBadRequest      = errors.New("agent: bad request")
	ErrAgentAPIInternalError   = errors.New("agent: internal error")
	ErrAgentAPIConflict        = errors.New("agent: conflict")
	ErrAgentAPIUnknownResponse = errors.New("agent: unknown response")
)
