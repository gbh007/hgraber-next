package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
)

var errImplementMe = errors.New("implement me")

func (c *Controller) APIAgentNewPost(ctx context.Context, req *server.APIAgentNewPostReq) (server.APIAgentNewPostRes, error) {
	return nil, errImplementMe
}

func (c *Controller) APIAgentTaskExportPost(ctx context.Context, req *server.APIAgentTaskExportPostReq) (server.APIAgentTaskExportPostRes, error) {
	return nil, errImplementMe
}

func (c *Controller) APIRatePost(ctx context.Context, req *server.APIRatePostReq) (server.APIRatePostRes, error) {
	return nil, errImplementMe
}

func (c *Controller) APIUserLoginPost(ctx context.Context, req *server.APIUserLoginPostReq) (server.APIUserLoginPostRes, error) {
	return nil, errImplementMe
}

func (c *Controller) APIUserRegistrationPost(ctx context.Context, req *server.APIUserRegistrationPostReq) (server.APIUserRegistrationPostRes, error) {
	return nil, errImplementMe
}
