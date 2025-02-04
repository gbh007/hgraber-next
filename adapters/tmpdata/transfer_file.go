package tmpdata

import (
	"github.com/gbh007/hgraber-next/entities"
)

func (s *Storage) AddToFileTransfer(transfers []entities.FileTransfer) {
	s.toFileTransfer.Push(transfers)
}

func (s *Storage) FileTransferList() []entities.FileTransfer {
	return s.toFileTransfer.Pop()
}
