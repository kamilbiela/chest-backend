package storage

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/dchest/uniuri"
)

type Local struct {
	basePath string
}

func (l *Local) dirpath(filename string) string {
	return l.basePath + "/" + strings.ToLower(filename[0:2])
}

func (l *Local) fullpath(filename string) string {
	return l.dirpath(filename) + "/" + filename
}

func (l *Local) createDir(filename string) error {
	return os.MkdirAll(l.dirpath(filename), 0770)
}

func (l *Local) Save(reader io.Reader) (string, error) {
	// @todo there is small chance of collision, add check for that
	filename := uniuri.NewLen(32)
	fullpath := l.fullpath(filename)

	err := l.createDir(filename)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(file, reader)

	if err != nil {
		log.Fatalln(err)
	}

	return filename, nil
}

func (l *Local) Get(filename string) (io.ReadCloser, error) {
	return os.Open(l.fullpath(filename))
}

func NewLocal(basePath string) *Local {
	return &Local{
		basePath: basePath,
	}
}
