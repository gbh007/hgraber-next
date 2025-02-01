package tmpdata

import (
	"hgnext/internal/entities"
)

func (s *Storage) AddToFileTransfer(transfers []entities.FileTransfer) {
	s.toFileTransfer.Push(transfers)
}

func (s *Storage) FileTransferList() []entities.FileTransfer {
	return s.toFileTransfer.Pop()
}
