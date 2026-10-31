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
	MasterAPIOffline         = errors.New("master: offline")
	MasterAPIUnauthorized    = errors.New("master: unauthorized")
	MasterAPIForbidden       = errors.New("master: forbidden")
	MasterAPIBadRequest      = errors.New("master: bad request")
	MasterAPIInternalError   = errors.New("master: internal error")
	MasterAPIConflict        = errors.New("master: conflict")
	MasterAPIUnknownResponse = errors.New("master: unknown response")
)
