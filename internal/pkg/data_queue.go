package pkg

import "sync"

type DataQueue[T any] struct {
	queue      []T
	mutex      *sync.Mutex
	bufferSize int
}

func NewDataQueue[T any](bufferSize int) *DataQueue[T] {
	return &DataQueue[T]{
		queue:      make([]T, 0, bufferSize),
		mutex:      &sync.Mutex{},
		bufferSize: bufferSize,
	}
}

func (dq *DataQueue[T]) Push(s []T) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.queue = append(dq.queue, s...)
}

func (dq *DataQueue[T]) PushOne(s T) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	dq.queue = append(dq.queue, s)
}

func (dq *DataQueue[T]) Pop() []T {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	if len(dq.queue) == 0 {
		return nil
	}

	s := dq.queue
	dq.queue = make([]T, 0, dq.bufferSize)

	return s
}

func (dq *DataQueue[T]) PopOne() (T, bool) {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	if len(dq.queue) == 0 {
		var empty T
		return empty, false
	}

	e := dq.queue[0]

	dq.queue = dq.queue[1:]

	return e, true
}

func (dq *DataQueue[T]) Size() int {
	dq.mutex.Lock()
	defer dq.mutex.Unlock()

	return len(dq.queue)
}
