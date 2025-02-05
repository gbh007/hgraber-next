package apiservercore

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func OptURL(u *url.URL) serverAPI.OptURI {
	if u == nil {
		return serverAPI.OptURI{}
	}

	return serverAPI.NewOptURI(*u)
}

func UrlFromOpt(u serverAPI.OptURI) *url.URL {
	if !u.Set {
		return nil
	}

	return &u.Value
}

func OptTime(t time.Time) serverAPI.OptDateTime {
	if t.IsZero() {
		return serverAPI.OptDateTime{}
	}

	return serverAPI.NewOptDateTime(t)
}

func OptString(s string) serverAPI.OptString {
	if s == "" {
		return serverAPI.OptString{}
	}

	return serverAPI.NewOptString(s)
}

func OptUUID(u uuid.UUID) serverAPI.OptUUID {
	if u == uuid.Nil {
		return serverAPI.OptUUID{}
	}

	return serverAPI.NewOptUUID(u)
}
