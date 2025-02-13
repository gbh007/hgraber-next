package tmpdata

import "github.com/gbh007/hgraber-next/domain/agentmodel"

func (s *Storage) AddToExport(books []agentmodel.BookToExport) {
	s.toExport.Push(books)
}

func (s *Storage) ExportList() []agentmodel.BookToExport {
	return s.toExport.Pop()
}
