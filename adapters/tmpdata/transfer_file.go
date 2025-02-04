package tmpdata

import (
	"github.com/gbh007/hgraber-next/domain/core"
)

func (s *Storage) AddToFileTransfer(transfers []core.FileTransfer) {
	s.toFileTransfer.Push(transfers)
}

func (s *Storage) FileTransferList() []core.FileTransfer {
	return s.toFileTransfer.Pop()
}
