package core

import "errors"

var (
	BookNotFoundError           = errors.New("book not found")
	PageNotFoundError           = errors.New("page not found")
	FileNotFoundError           = errors.New("file not found")
	AgentNotFoundError          = errors.New("agent not found")
	BookAlreadyExistsError      = errors.New("book already exists")
	AttributeRemapNotFoundError = errors.New("attribute remap not found")
	UnsupportedAttributeError   = errors.New("attribute is not supported")
	MissingFSError              = errors.New("missing fs")
)
