package entities

import "github.com/google/uuid"

type RebuildBookRequest struct {
	OldBook              BookFull
	SelectedPages        []int
	MergeWithBook        uuid.UUID
	OnlyUniquePages      bool
	ExcludeDeadHashPages bool
}
