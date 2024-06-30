package tmpdata

import "hgnext/internal/entities"

type Storage struct {
	toExport *dataQueue[entities.BookFullWithAgent]
}

func New() *Storage {
	return &Storage{
		toExport: newDataQueue[entities.BookFullWithAgent](100),
	}
}
