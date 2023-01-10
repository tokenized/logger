package logger

import (
	"io"
	"os"
	"sync"

	"github.com/pkg/errors"
)

func newOutput(path string) (io.Writer, error) {
	if len(path) > 0 {
		if path == "dummy" { // for benchmarking
			return &dummyWriter{}, nil
		} else {
			file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return nil, errors.Wrap(err, "open file")
			}

			return file, nil
			// return &fileWriter{file: file}, nil
		}
	} else {
		return os.Stderr, nil
		// return &stdErrWriter{}, nil
	}
}

type fileWriter struct {
	file *os.File
	lock sync.Mutex
}

func (w *fileWriter) Write(b []byte) (int, error) {
	w.lock.Lock()
	n, err := w.file.Write(b)
	// w.file.Sync()
	w.lock.Unlock()
	return n, err
}

type stdErrWriter struct {
	lock sync.Mutex
}

func (w *stdErrWriter) Write(b []byte) (int, error) {
	w.lock.Lock()
	n, err := os.Stderr.Write(b)
	// os.Stderr.Sync()
	w.lock.Unlock()
	return n, err
}

type dummyWriter struct{}

func (d *dummyWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
