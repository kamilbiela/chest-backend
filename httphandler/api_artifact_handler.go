package httphandler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamilbiela/chest-backend/lib"
)

func ApiGetArtifactsHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"data\": [1, 2, 3]}")
	})
}

func ApiGetArtifactHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		filename, ok := vars["filename"]

		if !ok {
			log.Println("Not found")
			return
		}

		readCloser, err := container.Storage().Get(filename)

		if err != nil {
			log.Println(err)
			return
		}

		defer readCloser.Close()

		if !ok {
			log.Println("Not found")
			return
		}

		io.Copy(w, readCloser)
	})
}

func ApiPostArtifactHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := container.Storage().Save("testfilename", r.Body)

		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}

		log.Println(path)

		w.WriteHeader(200)
		return
	})
}
