package apiserver

import (
	"context"
	"errors"
	"net/http"

	"hgnext/internal/controllers/apiserver/internal/server"
)

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t server.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) HandleCookies(ctx context.Context, operationName string, t server.Cookies) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}

func (c *Controller) APIUserLoginPost(ctx context.Context, req *server.APIUserLoginPostReq) (server.APIUserLoginPostRes, error) {
	cookie := http.Cookie{
		Name:     "X-HG-Token",
		Value:    req.Token,
		Path:     "/",
		HttpOnly: true,
	}

	return &server.APIUserLoginPostNoContent{
		SetCookie: server.NewOptString(cookie.String()),
	}, nil
}
