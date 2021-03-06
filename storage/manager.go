package storage

import (
	"io"
)

type Manager interface {
	// Saves passed file, can change name of file and return new name.
	Save(rc io.Reader) (string, error)

	// Get file by filename
	Get(filename string) (io.ReadCloser, error)

	// Delete file by filename
	Delete(filename string) error
}
