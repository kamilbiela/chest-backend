package main

import (
	"github.com/gorilla/pat"
	"github.com/turbogopher/chest-backend/httphandler"
	"log"
	"net/http"
)

func main() {
	container := NewContainer()

	r := pat.New()
	r.Get("/", httphandler.Index(container.Config()))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./../chest-frontend/")))

	http.Handle("/", r)

	log.Println("Listenig on" + container.Config().HTTPAddress)
	err := http.ListenAndServe(container.Config().HTTPAddress, r)
	if err != nil {
		log.Fatalln(err)
	}
}
