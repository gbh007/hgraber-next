package systemhandlers

import (
	"context"
	"slices"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *SystemHandlersController) APISystemTaskResultsGet(
	ctx context.Context,
) (serverapi.APISystemTaskResultsGetRes, error) {
	result, err := c.systemUseCases.TaskResults(ctx)
	if err != nil {
		return &serverapi.APISystemTaskResultsGetInternalServerError{
			InnerCode: apiservercore.TaskerUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	responseResults := pkg.Map(
		result,
		func(raw *systemmodel.TaskResult) serverapi.APISystemTaskResultsGetOKResultsItem {
			return serverapi.APISystemTaskResultsGetOKResultsItem{
				Name:              raw.Name,
				Error:             apiservercore.OptString(raw.Error),
				Result:            apiservercore.OptString(raw.Result),
				DurationFormatted: max(raw.Duration(), 0).String(),
				StartedAt:         raw.StartedAt,
				EndedAt:           raw.EndedAt,
				Stages: pkg.Map(
					raw.Stages,
					//nolint:lll // будет исправлено позднее
					func(rawStage *systemmodel.TaskResultStage) serverapi.APISystemTaskResultsGetOKResultsItemStagesItem {
						return serverapi.APISystemTaskResultsGetOKResultsItemStagesItem{
							Name:              rawStage.Name,
							Error:             apiservercore.OptString(rawStage.Error),
							Result:            apiservercore.OptString(rawStage.Result),
							DurationFormatted: max(rawStage.Duration(), 0).String(),
							StartedAt:         rawStage.StartedAt,
							EndedAt:           rawStage.EndedAt,
							Progress:          rawStage.Progress,
							Total:             rawStage.Total,
						}
					},
				),
			}
		},
	)

	slices.SortFunc(responseResults, func(a, b serverapi.APISystemTaskResultsGetOKResultsItem) int {
		return b.StartedAt.Compare(a.StartedAt)
	})

	//nolint:lll // текст для отдачи с сервера
	return &serverapi.APISystemTaskResultsGetOK{
		Results: responseResults,
		Tasks: []serverapi.APISystemTaskResultsGetOKTasksItem{
			{
				Code: "clean_after_rebuild",
				Name: "Очистить данные после ребилдов",
				Description: serverapi.NewOptString(
					"Наполняет мертвые хеши, очищает удаленные страницы, удаляет неиспользуемые файлы, удаляет удаленные ребилды",
				),
			},
			{
				Code: "clean_after_parse",
				Name: "Очистить данные после парсинга и закачки книг",
				Description: serverapi.NewOptString(
					"Перепривязывает страницы к одному общему файлу, вместо двух одинаковых файлов, удаляет неиспользуемые файлы",
				),
			},
			{
				Code: "deduplicate_files",
				Name: "Дедуплицировать файлы",
				Description: serverapi.NewOptString(
					"Перепривязывает страницы к одному общему файлу, вместо двух одинаковых файлов",
				),
			},
			{
				Code: "remove_detached_files",
				Name: "Удалить ни с чем не связанные файлы",
				Description: serverapi.NewOptString(
					"Удаляет из БД и файловой системы файлы, которые остались без привязки к страницам",
				),
			},
			{
				Code: "fill_dead_hashes",
				Name: "Наполнить мертвые хеши",
				Description: serverapi.NewOptString(
					"Наполняет мертвые хеши по данным удаленных страниц, для которых нет таких же хешей для \"живых\" страниц",
				),
			},
			{
				Code: "fill_dead_hashes_with_remove_deleted_pages",
				Name: "Наполнить мертвые хеши и удалить удаленные страницы с ними",
				Description: serverapi.NewOptString(
					"Наполняет мертвые хеши по данным удаленных страниц, для которых нет таких же хешей для \"живых\" страниц, после чего удаляет такие страницы из БД",
				),
			},
			{
				Code:        "clean_deleted_pages",
				Name:        "Очистить удаленные страницы",
				Description: serverapi.NewOptString("Очищает данные удаленных страниц"),
			},
			{
				Code: "clean_deleted_rebuilds",
				Name: "Очистить удаленные ребилды",
				Description: serverapi.NewOptString(
					"Удаляет полностью из БД данные ребилдов (страницы, метки, атрибуты и т.д.), что отмечены как удаленные",
				),
			},
			{
				Code:        "remap_attributes",
				Name:        "Выполнить ремапинг аттрибутов всех книг",
				Description: serverapi.NewOptString("Выполняет ремапинг аттрибутов всех книг, кроме удаленных"),
			},
		},
	}, nil
}
