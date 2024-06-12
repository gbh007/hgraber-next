package entities

import "errors"

var (
	BookNotFoundError         = errors.New("book not found")
	PageNotFoundError         = errors.New("page not found")
	FileNotFoundError         = errors.New("file not found")
	BookAlreadyExistsError    = errors.New("book already exists")
	UnsupportedAttributeError = errors.New("attribute is not supported")
)
