package apiserver

import (
	"context"
	"slices"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APITaskResultsGet(ctx context.Context) (serverAPI.APITaskResultsGetRes, error) {
	result, err := c.taskUseCases.TaskResults(ctx)
	if err != nil {
		return &serverAPI.APITaskResultsGetInternalServerError{
			InnerCode: TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	responseResults := pkg.Map(result, func(raw *entities.TaskResult) serverAPI.APITaskResultsGetOKResultsItem {
		return serverAPI.APITaskResultsGetOKResultsItem{
			Name:              raw.Name,
			Error:             optString(raw.Error),
			Result:            optString(raw.Result),
			DurationFormatted: max(raw.Duration(), 0).String(),
			StartedAt:         raw.StartedAt,
			EndedAt:           raw.EndedAt,
			Stages: pkg.Map(raw.Stages, func(rawStage *entities.TaskResultStage) serverAPI.APITaskResultsGetOKResultsItemStagesItem {
				return serverAPI.APITaskResultsGetOKResultsItemStagesItem{
					Name:              rawStage.Name,
					Error:             optString(rawStage.Error),
					Result:            optString(rawStage.Result),
					DurationFormatted: max(rawStage.Duration(), 0).String(),
					StartedAt:         rawStage.StartedAt,
					EndedAt:           rawStage.EndedAt,
					Progress:          rawStage.Progress,
					Total:             rawStage.Total,
				}
			}),
		}
	})

	slices.SortFunc(responseResults, func(a, b serverAPI.APITaskResultsGetOKResultsItem) int {
		return b.StartedAt.Compare(a.StartedAt)
	})

	return &serverAPI.APITaskResultsGetOK{
		Results: responseResults,
	}, nil
}
