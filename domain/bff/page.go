package bff

import (
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func PageWithHashToPreview(p core.PageWithHash) PreviewPage {
	return PreviewPage{
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		Downloaded: p.Downloaded,
		FileID:     p.FileID,
		FSID:       p.FSID,
	}
}

type PreviewPage struct {
	PageNumber  int
	Ext         string
	Downloaded  bool
	FileID      uuid.UUID
	FSID        uuid.UUID
	HasDeadHash StatusFlag
}
