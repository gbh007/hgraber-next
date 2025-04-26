package hproxymodel

import "net/url"

type ListBook struct {
	ExtURL        url.URL
	Name          string
	ExtPreviewURL *url.URL
}

type ListPage struct {
	ExtURL url.URL
	Name   string
}

type List struct {
	Books      []ListBook
	Pagination []ListPage
}

type BookPage struct {
	PageNumber    int
	ExtPreviewURL url.URL
}

type BookAttribute struct {
	Code   string
	Name   string
	Values []BookAttributeValue
}

type BookAttributeValue struct {
	ExtName string
	Name    string
	ExtURL  *url.URL
}

type Book struct {
	Name       string
	ExURL      url.URL
	PreviewURL *url.URL
	PageCount  int
	Pages      []BookPage
	Attributes []BookAttribute
}
