package entities

import "errors"

var (
	BookNotFoundError         = errors.New("book not found")
	PageNotFoundError         = errors.New("page not found")
	FileNotFoundError         = errors.New("file not found")
	AgentNotFoundError        = errors.New("agent not found")
	BookAlreadyExistsError    = errors.New("book already exists")
	UnsupportedAttributeError = errors.New("attribute is not supported")
)

var (
	AgentAPIOffline         = errors.New("agent: offline")
	AgentAPIUnauthorized    = errors.New("agent: unauthorized")
	AgentAPIForbidden       = errors.New("agent: forbidden")
	AgentAPIBadRequest      = errors.New("agent: bad request")
	AgentAPIInternalError   = errors.New("agent: internal error")
	AgentAPIConflict        = errors.New("agent: conflict")
	AgentAPIUnknownResponse = errors.New("agent: unknown response")
)
