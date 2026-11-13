package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/samber/lo"
)

func (c *LabelHandlersController) APILabelPresetListGet(
	ctx context.Context,
) (*serverapi.APILabelPresetListGetOK, error) {
	presets, err := c.labelUseCases.LabelPresets(ctx)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APILabelPresetListGetOK{
		Presets: lo.Map(presets, func(raw core.BookLabelPreset, _ int) serverapi.APILabelPresetListGetOKPresetsItem {
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
