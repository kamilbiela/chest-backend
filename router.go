package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/kamilbiela/chest-backend/httphandler"
	"github.com/kamilbiela/chest-backend/lib"
)

func setupRouter(container *lib.Container, r *mux.Router) *mux.Router {
	defaultChain := alice.New(container.Auth().Middleware)

	r.Handle("/login/github", httphandler.LoginGithubHandler(container))

	r.Handle("/login/github/authorized", httphandler.LoginGithubAcceptedHandler(container))

	r.Handle("/api/project", defaultChain.Then(httphandler.ApiGetProjectsHandler(container))).
		Methods("GET")

	r.Handle("/api/setting", defaultChain.Then(httphandler.ApiGetSettingsHandler(container))).
		Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../chest-frontend/")))

	return r
}
