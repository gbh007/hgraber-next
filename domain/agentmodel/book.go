package agentmodel

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type BookWithAgent struct {
	core.Book

	AgentID uuid.UUID
}

type BookToExport struct {
	AgentID uuid.UUID
	BookID  uuid.UUID

	DeleteAfterExport bool
}
