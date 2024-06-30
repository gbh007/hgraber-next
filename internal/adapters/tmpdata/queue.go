package tmpdata

import "sync"

type dataQueue[T any] struct {
	queue      []T
	mutex      *sync.Mutex
	bufferSize int
}

func newDataQueue[T any](bufferSize int) *dataQueue[T] {
	return &dataQueue[T]{
		queue:      make([]T, 0, bufferSize),
		mutex:      &sync.Mutex{},
		bufferSize: bufferSize,
	}
}

func (dq *dataQueue[T]) Push(s []T) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.queue = append(dq.queue, s...)
}

func (dq *dataQueue[T]) Pop() []T {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	s := dq.queue
	dq.queue = make([]T, 0, dq.bufferSize)

	return s
}
