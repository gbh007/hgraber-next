package systemhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *SystemHandlersController) APISystemWorkerConfigPost(ctx context.Context, req *serverapi.APISystemWorkerConfigPostReq) (serverapi.APISystemWorkerConfigPostRes, error) {
	counts := make(map[string]int, len(req.RunnersCount))

	for _, runnerCount := range req.RunnersCount {
		counts[runnerCount.Name] = runnerCount.Count
	}

	c.systemUseCases.SetWorkerConfig(ctx, counts)

	return &serverapi.APISystemWorkerConfigPostNoContent{}, nil
}
