package entities

import "errors"

var (
	BookNotFoundError         = errors.New("book not found")
	PageNotFoundError         = errors.New("page not found")
	FileNotFoundError         = errors.New("file not found")
	FileAlreadyExistsError    = errors.New("file already exists")
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

var (
	MasterAPIOffline         = errors.New("master: offline")
	MasterAPIUnauthorized    = errors.New("master: unauthorized")
	MasterAPIForbidden       = errors.New("master: forbidden")
	MasterAPIBadRequest      = errors.New("master: bad request")
	MasterAPIInternalError   = errors.New("master: internal error")
	MasterAPIConflict        = errors.New("master: conflict")
	MasterAPIUnknownResponse = errors.New("master: unknown response")
)
