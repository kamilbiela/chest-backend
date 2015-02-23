package httphandler

import (
	"net/http"

	"github.com/kamilbiela/chest-backend/lib"
	"github.com/kamilbiela/chest-backend/model"
)

type ProjectsResponse struct {
	Projects []*model.Project
}

func ApiGetProjectsHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
