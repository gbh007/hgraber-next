package serveragentdto

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

var (
	ErrMasterAPIOffline         = errors.New("master: offline")
	ErrMasterAPIUnauthorized    = errors.New("master: unauthorized")
	ErrMasterAPIForbidden       = errors.New("master: forbidden")
	ErrMasterAPIBadRequest      = errors.New("master: bad request")
	ErrMasterAPIInternalError   = errors.New("master: internal error")
	ErrMasterAPIConflict        = errors.New("master: conflict")
	ErrMasterAPIUnknownResponse = errors.New("master: unknown response")
)
