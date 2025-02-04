package agentmodel

import "errors"

var (
	AgentAPIOffline         = errors.New("agent: offline")
	AgentAPIUnauthorized    = errors.New("agent: unauthorized")
	AgentAPIForbidden       = errors.New("agent: forbidden")
	AgentAPIBadRequest      = errors.New("agent: bad request")
	AgentAPIInternalError   = errors.New("agent: internal error")
	AgentAPIConflict        = errors.New("agent: conflict")
	AgentAPIUnknownResponse = errors.New("agent: unknown response")
)
