package apiserver

import (
	"context"

	"hgnext/open_api/serverAPI"
)

func (c *Controller) APISystemWorkerConfigPost(ctx context.Context, req *serverAPI.APISystemWorkerConfigPostReq) (serverAPI.APISystemWorkerConfigPostRes, error) {
	counts := make(map[string]int, len(req.RunnersCount))

	for _, runnerCount := range req.RunnersCount {
		counts[runnerCount.Name] = runnerCount.Count
	}

	c.webAPIUseCases.SetWorkerConfig(ctx, counts)

	return &serverAPI.APISystemWorkerConfigPostNoContent{}, nil
}
