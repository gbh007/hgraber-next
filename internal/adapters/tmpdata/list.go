package tmpdata

import "sync"

type dataList[T any] struct {
	list       []T
	mutex      sync.RWMutex
	bufferSize int
}

func newDataList[T any](bufferSize int) *dataList[T] {
	return &dataList[T]{
		list:       make([]T, 0, bufferSize),
		bufferSize: bufferSize,
	}
}

func (dq *dataList[T]) Push(s []T) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.list = append(dq.list, s...)
}

func (dq *dataList[T]) PushOne(s T) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.list = append(dq.list, s)
}

func (dq *dataList[T]) All() []T {
	dq.mutex.RLock()
	defer dq.mutex.RUnlock()

	s := dq.list

	return s
}
