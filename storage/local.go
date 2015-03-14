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
	// get unique filename
	var filename, fullpath string
	var err error
	var file *os.File
	for tries := 0; ; tries++ {
		if tries > 10 {
			panic("Too many tries to save a file")
		}
		filename = strings.ToLower(uniuri.NewLen(uniuri.UUIDLen))
		fullpath = l.fullpath(filename)

		err = l.createDir(filename)
		if err != nil {
			log.Fatalln(err)
		}

		file, err = os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		defer file.Close()

		if err == nil {
			break
		} else {
			if e, ok := err.(*os.PathError); ok && os.IsNotExist(e.Err) {
				// @todo log/alert this
				return "", err
			}
		}
	}

	_, err = io.Copy(file, reader)

	if err != nil {
		return "", err
	}

	return filename, nil
}

func (l *Local) Get(filename string) (io.ReadCloser, error) {
	return os.Open(l.fullpath(filename))
}

func (l *Local) Delete(filename string) error {
	return os.Remove(l.fullpath(filename))
}

func NewLocal(basePath string) *Local {
	return &Local{
		basePath: basePath,
	}
}
