//nolint:decorder // будет исправлено позднее
package core

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type BookCalc struct {
	CalcPageCount     *int64
	CalcFileCount     *int64
	CalcDeadHashCount *int64
	CalcPageSize      *int64
	CalcFileSize      *int64
	CalcDeadHashSize  *int64
	CalculatedAt      time.Time
}

type Book struct {
	ID               uuid.UUID
	Name             string
	OriginURL        *url.URL
	PageCount        int
	AttributesParsed bool
	CreateAt         time.Time

	Deleted    bool
	DeletedAt  time.Time
	Verified   bool
	VerifiedAt time.Time

	IsRebuild bool

	Calc BookCalc
}

func (b Book) IsLoaded() bool {
	return !b.Deleted && b.AttributesParsed && b.PageCount > 0 && b.Name != ""
}

func (b Book) ParsedName() bool {
	return b.Name != ""
}

// FIXME: подумать что делать с такими моделями
type BookContainer struct {
	Book       Book
	Pages      []Page
	Attributes map[string][]string
	Labels     []BookLabel
}

func (b BookContainer) IsLoaded() bool {
	if !b.Book.IsLoaded() {
		return false
	}

	for _, p := range b.Pages {
		if !p.IsLoaded() {
			return false
		}
	}

	return true
}

func (b BookContainer) Filename() string {
	return b.Book.ID.String() + " " + EscapeBookFileName(b.Book.Name) + ".zip"
}

type BookSize struct {
	Unique                  int64
	UniqueWithoutDeadHashes int64
	Shared                  int64
	DeadHashes              int64
	Total                   int64

	// TODO: технически это не размер, лучше всего будет отделить
	UniqueCount                  int
	UniqueWithoutDeadHashesCount int
	SharedCount                  int
	DeadHashesCount              int
	InnerDuplicateCount          int
}
