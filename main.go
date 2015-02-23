package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamilbiela/chest-backend/lib"
)

var isSetup bool

func init() {
	flag.BoolVar(&isSetup, "setup", false, "--setup to run app setup (insert sql into db)")
	flag.Parse()
}

func main() {
	container := lib.NewContainer()

	log.Println(isSetup)

	if isSetup {
		setup(container)
	} else {
		startApp(container)
	}
}

func startApp(container *lib.Container) {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db := container.Database()
	err = db.Ping()
	if err != nil {
		log.Println("Can't connect to database")
		log.Fatalln(err)
	}

	cache := container.Cache()
	err = cache.Ping()
	if err != nil {
		log.Println("Can't connect to " + container.Config().Cache.Strategy)
		log.Fatalln(err)
	}

	r := setupRouter(container, mux.NewRouter())
	http.Handle("/", r)

	log.Println("Listenig on" + container.Config().HTTPAddress)
	err = http.ListenAndServe(container.Config().HTTPAddress, r)
	if err != nil {
		log.Fatalln(err)
	}
}
