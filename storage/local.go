package storage

import (
	"io"
	"log"
	"os"
)

type Local struct {
	basePath string
}

func (l *Local) fullpath(filename string) string {
	// @todo clean filename
	return l.basePath + "/" + filename
}

func (l *Local) Save(filename string, rc io.ReadCloser) (string, error) {
	fullpath := l.fullpath(filename)
	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(file, rc)
	defer rc.Close()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Saved file to: " + fullpath)

	return "", nil
}

func (l *Local) Get(filename string) (io.ReadCloser, error) {
	return os.Open(l.fullpath(filename))
}

func NewLocal(basePath string) *Local {
	return &Local{
		basePath: basePath,
	}
}
