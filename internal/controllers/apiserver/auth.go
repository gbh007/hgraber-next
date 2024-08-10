package apiserver

import (
	"context"
	"errors"
	"net/http"

	"hgnext/open_api/serverAPI"
)

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t serverAPI.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) HandleCookies(ctx context.Context, operationName string, t serverAPI.Cookies) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) APIUserLoginPost(ctx context.Context, req *serverAPI.APIUserLoginPostReq) (serverAPI.APIUserLoginPostRes, error) {
	cookie := http.Cookie{
		Name:     "X-HG-Token",
		Value:    req.Token,
		Path:     "/",
		HttpOnly: true,
	}

	return &serverAPI.APIUserLoginPostNoContent{
		SetCookie: serverAPI.NewOptString(cookie.String()),
	}, nil
}
