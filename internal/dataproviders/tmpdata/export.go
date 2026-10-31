package tmpdata

import "github.com/gbh007/hgraber-next/internal/domain/serveragentdto"

func (s *Storage) AddToExport(books []serveragentdto.BookToExport) {
	s.toExport.Push(books)
}

func (s *Storage) ExportList() []serveragentdto.BookToExport {
	return s.toExport.Pop()
}
