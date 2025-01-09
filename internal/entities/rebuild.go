package entities

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrRebuildBookForbiddenMerge    = errors.New("merge with book forbidden")
	ErrRebuildBookIncorrectRequest  = errors.New("incorrect request")
	ErrRebuildBookEmptyPages        = errors.New("empty pages on rebuild")
	ErrRebuildBookMissingSourcePage = errors.New("missing source page")
)

type RebuildBookRequest struct {
	ModifiedOldBook BookFull
	SelectedPages   []int
	MergeWithBook   uuid.UUID

	Flags RebuildBookRequestFlags
}

type RebuildBookRequestFlags struct {
	OnlyUniquePages      bool
	ExcludeDeadHashPages bool
	Only1CopyPages       bool

	SetOriginLabels bool
	AutoVerify      bool

	ExtractMode bool

	MarkUnusedPagesAsDeadHash              bool
	MarkUnusedPagesAsDeleted               bool
	MarkEmptyBookAsDeletedAfterRemovePages bool
}
