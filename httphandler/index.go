package httphandler

import (
	"github.com/turbogopher/chest-backend/lib"
	"net/http"
)

func Index(config *lib.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, config.WWWDir+"/index.html")
	}
}
