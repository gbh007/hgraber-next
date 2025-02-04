package apiserver

import (
	"context"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APILabelPresetListGet(ctx context.Context) (serverAPI.APILabelPresetListGetRes, error) {
	presets, err := c.webAPIUseCases.LabelPresets(ctx)
	if err != nil {
		return &serverAPI.APILabelPresetListGetInternalServerError{
			InnerCode: WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetListGetOK{
		Presets: pkg.Map(presets, func(raw entities.BookLabelPreset) serverAPI.APILabelPresetListGetOKPresetsItem {
			return serverAPI.APILabelPresetListGetOKPresetsItem{
				Name:        raw.Name,
				Description: optString(raw.Description),
				Values:      raw.Values,
				CreatedAt:   raw.CreatedAt,
				UpdatedAt:   optTime(raw.UpdatedAt),
			}
		}),
	}, nil
}
