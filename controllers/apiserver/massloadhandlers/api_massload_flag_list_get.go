package massloadhandlers

import (
	"context"
	"time"

	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *MassloadController) APIMassloadFlagListGet(ctx context.Context) (serverapi.APIMassloadFlagListGetRes, error) {
	return &serverapi.APIMassloadFlagListGetOK{
		Flags: []serverapi.APIMassloadFlagListGetOKFlagsItem{
			{
				Code:      "deduplicated",
				Name:      "Дедуплицирована",
				CreatedAt: time.Now(),
			},
		},
	}, nil
}
