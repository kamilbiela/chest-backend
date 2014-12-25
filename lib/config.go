package lib

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	WWWDir      string
	HTTPAddress string
}

func NewConfig() *Config {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dir)

	return &Config{
		WWWDir:      dir + "/../chest-frontend",
		HTTPAddress: ":3000",
	}
}
