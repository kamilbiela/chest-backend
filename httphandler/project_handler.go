package httphandler

import (
	"fmt"
	"net/http"

	"github.com/kamilbiela/chest-backend/lib"
)

func ProjectHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "{\"data\": [1, 2, 3]}")

	})
}
