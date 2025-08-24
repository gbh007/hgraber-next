package apiserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(
	ctx context.Context,
	operationName string,
	t serverapi.HeaderAuth,
) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) HandleCookies(
	ctx context.Context,
	operationName string,
	t serverapi.Cookies,
) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) APIUserLoginPost(
	ctx context.Context,
	req *serverapi.APIUserLoginPostReq,
) (serverapi.APIUserLoginPostRes, error) {
	cookie := http.Cookie{
		Name:     "X-HG-Token",
		Value:    req.Token,
		Path:     "/",
		HttpOnly: true,
	}

	return &serverapi.APIUserLoginPostNoContent{
		SetCookie: serverapi.NewOptString(cookie.String()),
	}, nil
}
