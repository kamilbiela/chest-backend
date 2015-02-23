package lib

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	WWWDir      string
	HTTPAddress string
	Github      struct {
		ClientID             string
		ClientSecret         string
		AllowedOrganizations []string `xml:",any"`
	}
	MySQL struct {
		DSN string
	}
	Cache struct {
		Strategy string
		Memcache struct {
			Hosts []string `xml:",any"`
		}
	}
	Secret string
}

func NewConfig() *Config {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	config := Config{
		WWWDir:      dir + "/../chest-frontend",
		HTTPAddress: ":3000",
	}

	data, err := ioutil.ReadFile(dir + "/config.xml")
	if err != nil {
		log.Fatalln(err)
	}

	err = xml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err)
	}

	config.WWWDir = dir + config.WWWDir

	log.Println(config)

	return &config
}
