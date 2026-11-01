package parserdto

import (
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/core"
)

const (
	AttributeCodeAuthor    = core.AttributeCodeAuthor
	AttributeCodeCategory  = core.AttributeCodeCategory
	AttributeCodeCharacter = core.AttributeCodeCharacter
	AttributeCodeGroup     = core.AttributeCodeGroup
	AttributeCodeLanguage  = core.AttributeCodeLanguage
	AttributeCodeParody    = core.AttributeCodeParody
	AttributeCodeTag       = core.AttributeCodeTag
)

var AllAttributes = core.AllAttributes

type Book struct {
	URL        url.URL
	PreviewURL *url.URL

	Name string

	PageCount int
	Pages     []Page

	Attributes []Attribute
}

type Attribute struct {
	Code   string
	Values []AttributeValue
}

type AttributeValue struct {
	Value string
	URL   *url.URL
}

type Page struct {
	PageNumber int
	URL        url.URL
	Filename   string
}

type BookListUnitType byte

const (
	UnknownHProxyListUnitType BookListUnitType = iota
	ListHProxyListUnitType
	DetailsHProxyListUnitType
)

type BookList struct {
	Units   []BookListUnit
	NextURL *url.URL
}

type BookListUnit struct {
	LinkURL    url.URL
	PreviewURL *url.URL
	Name       string
	Type       BookListUnitType
}
