package labelhandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *LabelHandlersController) APILabelPresetListGet(
	ctx context.Context,
) (*serverapi.APILabelPresetListGetOK, error) {
	presets, err := c.labelUseCases.LabelPresets(ctx)
	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.LabelUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &serverapi.APILabelPresetListGetOK{
		Presets: pkg.Map(presets, func(raw core.BookLabelPreset) serverapi.APILabelPresetListGetOKPresetsItem {
			return serverapi.APILabelPresetListGetOKPresetsItem{
				Name:        raw.Name,
				Description: apiservercore.OptString(raw.Description),
				Values:      raw.Values,
				CreatedAt:   raw.CreatedAt,
				UpdatedAt:   apiservercore.OptTime(raw.UpdatedAt),
			}
		}),
	}, nil
}
