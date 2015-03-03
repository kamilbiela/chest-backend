package storage

import (
	"io"
)

type Manager interface {
	Save(filename string, rc io.ReadCloser) (string, error)
	Get(filename string) (io.ReadCloser, error)
}
