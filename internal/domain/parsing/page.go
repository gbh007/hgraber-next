package parsing

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/domain/core"
)

type PageForDownloadWithAgent struct {
	core.PageForDownload

	AgentID uuid.UUID
}
