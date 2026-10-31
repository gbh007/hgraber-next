package parserdto

import "net/url"

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
