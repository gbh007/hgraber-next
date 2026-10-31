package api

import (
	"net/url"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func OptURL(u *url.URL) agentapi.OptURI {
	if u == nil {
		return agentapi.OptURI{}
	}

	return agentapi.NewOptURI(*u)
}

func UrlFromOpt(u agentapi.OptURI) *url.URL {
	if !u.Set {
		return nil
	}

	return &u.Value
}

func OptString(s string) agentapi.OptString {
	if s == "" {
		return agentapi.OptString{}
	}

	return agentapi.NewOptString(s)
}
