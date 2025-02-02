package tmpdata

import (
	"github.com/gbh007/hgraber-next/internal/entities"
)

func (s *Storage) AddToExport(books []entities.BookFullWithAgent) {
	s.toExport.Push(books)
}

func (s *Storage) ExportList() []entities.BookFullWithAgent {
	return s.toExport.Pop()
}
