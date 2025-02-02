package tmpdata

import (
	"github.com/gbh007/hgraber-next/internal/entities"
)

func (s *Storage) AddToFileTransfer(transfers []entities.FileTransfer) {
	s.toFileTransfer.Push(transfers)
}

func (s *Storage) FileTransferList() []entities.FileTransfer {
	return s.toFileTransfer.Pop()
}
