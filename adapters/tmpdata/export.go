package tmpdata

import "github.com/gbh007/hgraber-next/domain/agentmodel"

func (s *Storage) AddToExport(books []agentmodel.BookFullWithAgent) {
	s.toExport.Push(books)
}

func (s *Storage) ExportList() []agentmodel.BookFullWithAgent {
	return s.toExport.Pop()
}
