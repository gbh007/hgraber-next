package entities

import "net/url"

type HProxyBookDetails struct {
	URL        url.URL
	PreviewURL *url.URL

	Name string

	PageCount int
	Pages     []HProxyBookDetailsPagesItem

	Attributes []HProxyBookDetailsAttributesItem
}

type HProxyBookDetailsAttributesItem struct {
	Code   string
	Values []HProxyBookDetailsAttributesValueItem
}

type HProxyBookDetailsAttributesValueItem struct {
	Value string
	URL   *url.URL
}

type HProxyBookDetailsPagesItem struct {
	PageNumber int
	URL        url.URL
	Filename   string
}

type HProxyListUnitType byte

const (
	UnknownHProxyListUnitType HProxyListUnitType = iota
	ListHProxyListUnitType
	DetailsHProxyListUnitType
)

type HProxyList struct {
	Units   []HProxyListUnit
	NextURL *url.URL
}

type HProxyListUnit struct {
	LinkURL    url.URL
	PreviewURL *url.URL
	Name       string
	Type       HProxyListUnitType
}
