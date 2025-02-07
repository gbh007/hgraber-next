package tmpdata

import "github.com/gbh007/hgraber-next/domain/fsmodel"

func (s *Storage) AddToFileTransfer(transfers []fsmodel.FileTransfer) {
	s.toFileTransfer.Push(transfers)
}

func (s *Storage) FileTransferList() []fsmodel.FileTransfer {
	return s.toFileTransfer.Pop()
}
