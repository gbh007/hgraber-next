package mcp

import "github.com/google/uuid"

type bookData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	OriginURL  string    `json:"origin_url,omitempty"`
	PageCount  int       `json:"page_count"`
	Tags       []string  `json:"tags,omitempty"`
	Authors    []string  `json:"authors,omitempty"`
	Categories []string  `json:"categories,omitempty"`
	Characters []string  `json:"characters,omitempty"`
	Groups     []string  `json:"groups,omitempty"`
	Languages  []string  `json:"languages,omitempty"`
	Parodies   []string  `json:"parodies,omitempty"`
}

type attributeData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Count int    `json:"count"`
}

type hProxyAttributeValue struct {
	Value     string `json:"value"`
	OriginURL string `json:"origin_url,omitempty"`
}

type hProxyBookData struct {
	Name       string                 `json:"name"`
	PageCount  int                    `json:"page_count"`
	SystemIDs  []uuid.UUID            `json:"system_ids,omitempty"`
	OriginURL  string                 `json:"origin_url,omitempty"`
	Tags       []hProxyAttributeValue `json:"tags,omitempty"`
	Authors    []hProxyAttributeValue `json:"authors,omitempty"`
	Categories []hProxyAttributeValue `json:"categories,omitempty"`
	Characters []hProxyAttributeValue `json:"characters,omitempty"`
	Groups     []hProxyAttributeValue `json:"groups,omitempty"`
	Languages  []hProxyAttributeValue `json:"languages,omitempty"`
	Parodies   []hProxyAttributeValue `json:"parodies,omitempty"`
}
