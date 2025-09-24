package apiservercore

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func OptURL(u *url.URL) serverapi.OptURI {
	if u == nil {
		return serverapi.OptURI{}
	}

	return serverapi.NewOptURI(*u)
}

func URLFromOpt(u serverapi.OptURI) *url.URL {
	if !u.Set {
		return nil
	}

	return &u.Value
}

func OptTime(t time.Time) serverapi.OptDateTime {
	if t.IsZero() {
		return serverapi.OptDateTime{}
	}

	return serverapi.NewOptDateTime(t)
}

func OptString(s string) serverapi.OptString {
	if s == "" {
		return serverapi.OptString{}
	}

	return serverapi.NewOptString(s)
}

func OptInt64Pointer(v *int64) serverapi.OptInt64 {
	if v == nil {
		return serverapi.OptInt64{}
	}

	return serverapi.NewOptInt64(*v)
}

func OptUUID(u uuid.UUID) serverapi.OptUUID {
	if u == uuid.Nil {
		return serverapi.OptUUID{}
	}

	return serverapi.NewOptUUID(u)
}
