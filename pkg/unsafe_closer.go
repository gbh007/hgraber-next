package pkg

import (
	"errors"
	"io"
)

// unsafeCloser оболочка для закрытия потока, после полного чтения.
// Является костылем, отказаться с реализацией https://github.com/ogen-go/ogen/issues/1023.
type unsafeCloser struct {
	isClosed bool

	Body io.ReadCloser
}

func UnsafeCloser(body io.ReadCloser) io.Reader {
	return &unsafeCloser{
		Body: body,
	}
}

func (c *unsafeCloser) Read(p []byte) (n int, err error) {
	if c.isClosed || c.Body == nil {
		return 0, io.EOF
	}

	n, err = c.Body.Read(p)

	// FIXME: не только эта ошибка может возникнуть, что может привести к утечке дескриптора.
	// В новых версиях добавилась runtime.AddCleanup, по возможности перейти на нее.
	if errors.Is(err, io.EOF) {
		_ = c.Body.Close()
		c.isClosed = true
	}

	return n, err //nolint:wrapcheck // оставляем оригинальную ошибку
}
