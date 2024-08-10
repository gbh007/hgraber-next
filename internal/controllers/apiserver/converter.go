package apiserver

import (
	"net/url"
	"time"

	"hgnext/internal/controllers/apiserver/internal/server"
)

func optURL(u *url.URL) server.OptURI {
	if u == nil {
		return server.OptURI{}
	}

	return server.NewOptURI(*u)
}

func optTime(t time.Time) server.OptDateTime {
	if t.IsZero() {
		return server.OptDateTime{}
	}

	return server.NewOptDateTime(t)
}
