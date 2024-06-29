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
