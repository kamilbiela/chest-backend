package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/kamilbiela/chest-backend/httphandler"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	auth := container.Auth()

	r := mux.NewRouter()

	defaultChain := alice.New(auth.Middleware)

	r.Handle("/login/github", httphandler.LoginGithubHandler(container))
	r.Handle("/login/github/authorized", httphandler.LoginGithubAcceptedHandler(container))
	r.Handle("/api/project", defaultChain.Then(httphandler.ProjectHandler(container)))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../chest-frontend/")))

	http.Handle("/", r)

	log.Println("Listenig on" + container.Config().HTTPAddress)
	err := http.ListenAndServe(container.Config().HTTPAddress, r)
	if err != nil {
		log.Fatalln(err)
	}
}
