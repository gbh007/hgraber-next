package core

import "errors"

var (
	ErrBookNotFound           = errors.New("book not found")
	ErrPageNotFound           = errors.New("page not found")
	ErrFileNotFound           = errors.New("file not found")
	ErrAgentNotFound          = errors.New("agent not found")
	ErrBookAlreadyExists      = errors.New("book already exists")
	ErrAttributeRemapNotFound = errors.New("attribute remap not found")
	ErrUnsupportedAttribute   = errors.New("attribute is not supported")
	ErrMissingFS              = errors.New("missing fs")
)
