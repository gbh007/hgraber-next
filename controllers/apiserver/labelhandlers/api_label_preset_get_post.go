package labelhandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *LabelHandlersController) APILabelPresetGetPost(
	ctx context.Context,
	req *serverapi.APILabelPresetGetPostReq,
) (*serverapi.APILabelPresetGetPostOK, error) {
	raw, err := c.labelUseCases.LabelPreset(ctx, req.Name)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &serverapi.APILabelPresetGetPostOK{
		Name:        raw.Name,
		Description: apiservercore.OptString(raw.Description),
		Values:      raw.Values,
		CreatedAt:   raw.CreatedAt,
		UpdatedAt:   apiservercore.OptTime(raw.UpdatedAt),
	}, nil
}
