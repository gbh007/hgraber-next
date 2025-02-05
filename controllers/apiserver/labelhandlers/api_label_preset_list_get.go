package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *LabelHandlersController) APILabelPresetListGet(ctx context.Context) (serverAPI.APILabelPresetListGetRes, error) {
	presets, err := c.webAPIUseCases.LabelPresets(ctx)
	if err != nil {
		return &serverAPI.APILabelPresetListGetInternalServerError{
			InnerCode: apiservercore.WebAPIUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APILabelPresetListGetOK{
		Presets: pkg.Map(presets, func(raw core.BookLabelPreset) serverAPI.APILabelPresetListGetOKPresetsItem {
			return serverAPI.APILabelPresetListGetOKPresetsItem{
				Name:        raw.Name,
				Description: apiservercore.OptString(raw.Description),
				Values:      raw.Values,
				CreatedAt:   raw.CreatedAt,
				UpdatedAt:   apiservercore.OptTime(raw.UpdatedAt),
			}
		}),
	}, nil
}
