package entities

import "github.com/google/uuid"

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

	MarkUnusedPagesAsDeadHash              bool
	MarkUnusedPagesAsDeleted               bool
	MarkEmptyBookAsDeletedAfterRemovePages bool
}
