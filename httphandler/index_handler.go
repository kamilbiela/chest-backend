package httphandler

import (
	"net/http"

	"github.com/kamilbiela/chest-backend/lib"
)

func IndexHandler(config *lib.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, config.WWWDir+"/index.html")
	})
}
