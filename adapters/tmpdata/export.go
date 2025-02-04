package tmpdata

import (
	"github.com/gbh007/hgraber-next/domain/core"
)

func (s *Storage) AddToExport(books []core.BookFullWithAgent) {
	s.toExport.Push(books)
}

func (s *Storage) ExportList() []core.BookFullWithAgent {
	return s.toExport.Pop()
}
