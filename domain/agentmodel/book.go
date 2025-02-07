package agentmodel

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type BookWithAgent struct {
	core.Book
	AgentID uuid.UUID
}

type BookFullWithAgent struct {
	core.BookContainer
	AgentID uuid.UUID

	DeleteAfterExport bool
}
