package systemhandlers

import (
	"context"
	"slices"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APISystemTaskResultsGet(ctx context.Context) (serverAPI.APISystemTaskResultsGetRes, error) {
	result, err := c.taskUseCases.TaskResults(ctx)
	if err != nil {
		return &serverAPI.APISystemTaskResultsGetInternalServerError{
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	responseResults := pkg.Map(result, func(raw *core.TaskResult) serverAPI.APISystemTaskResultsGetOKResultsItem {
		return serverAPI.APISystemTaskResultsGetOKResultsItem{
			Name:              raw.Name,
			Error:             apiservercore.OptString(raw.Error),
			Result:            apiservercore.OptString(raw.Result),
			DurationFormatted: max(raw.Duration(), 0).String(),
			StartedAt:         raw.StartedAt,
			EndedAt:           raw.EndedAt,
			Stages: pkg.Map(raw.Stages, func(rawStage *core.TaskResultStage) serverAPI.APISystemTaskResultsGetOKResultsItemStagesItem {
				return serverAPI.APISystemTaskResultsGetOKResultsItemStagesItem{
					Name:              rawStage.Name,
					Error:             apiservercore.OptString(rawStage.Error),
					Result:            apiservercore.OptString(rawStage.Result),
					DurationFormatted: max(rawStage.Duration(), 0).String(),
					StartedAt:         rawStage.StartedAt,
					EndedAt:           rawStage.EndedAt,
					Progress:          rawStage.Progress,
					Total:             rawStage.Total,
				}
			}),
		}
	})

	slices.SortFunc(responseResults, func(a, b serverAPI.APISystemTaskResultsGetOKResultsItem) int {
		return b.StartedAt.Compare(a.StartedAt)
	})

	return &serverAPI.APISystemTaskResultsGetOK{
		Results: responseResults,
		Tasks: []serverAPI.APISystemTaskResultsGetOKTasksItem{
			{
				Code:        "deduplicate_files",
				Name:        "Дедуплицировать файлы",
				Description: serverAPI.NewOptString("Перепривязывает страницы к одному общему файлу, вместо двух одинаковых файлов"),
			},
			{
				Code:        "remove_detached_files",
				Name:        "Удалить ни с чем не связанные файлы",
				Description: serverAPI.NewOptString("Удаляет из БД и файловой системы файлы, которые остались без привязки к страницам"),
			},
			{
				Code:        "fill_dead_hashes",
				Name:        "Наполнить мертвые хеши",
				Description: serverAPI.NewOptString("Наполняет мертвые хеши по данным удаленных страниц, для которых нет таких же хешей для \"живых\" страниц"),
			},
			{
				Code:        "fill_dead_hashes_with_remove_deleted_pages",
				Name:        "Наполнить мертвые хеши и удалить удаленные страницы с ними",
				Description: serverAPI.NewOptString("Наполняет мертвые хеши по данным удаленных страниц, для которых нет таких же хешей для \"живых\" страниц, после чего удаляет такие страницы из БД"),
			},
			{
				Code:        "clean_deleted_pages",
				Name:        "Очистить удаленные страницы",
				Description: serverAPI.NewOptString("Очищает данные удаленных страниц"),
			},
			{
				Code:        "clean_deleted_rebuilds",
				Name:        "Очистить удаленные ребилды",
				Description: serverAPI.NewOptString("Удаляет полностью из БД данные ребилдов (страницы, метки, атрибуты и т.д.), что отмечены как удаленные"),
			},
		},
	}, nil
}
