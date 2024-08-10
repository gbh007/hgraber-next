package apiserver

import (
	"net/url"
	"time"

	"hgnext/open_api/serverAPI"
)

func optURL(u *url.URL) serverAPI.OptURI {
	if u == nil {
		return serverAPI.OptURI{}
	}

	return serverAPI.NewOptURI(*u)
}

func optTime(t time.Time) serverAPI.OptDateTime {
	if t.IsZero() {
		return serverAPI.OptDateTime{}
	}

	return serverAPI.NewOptDateTime(t)
}
