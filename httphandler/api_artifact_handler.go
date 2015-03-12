package httphandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dchest/uniuri"
	"github.com/gorilla/mux"
	"github.com/kamilbiela/chest-backend/lib"
	"github.com/nfnt/resize"
)

type ApiPostArtifactResponse struct {
	FileId      string
	ThumbFileId string
}

func ApiGetArtifactsHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"data\": [1, 2, 3]}")
	})
}

func ApiGetArtifactHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		filename, ok := params["filename"]

		if !ok {
			log.Println("Not found")
			return
		}

		readCloser, err := container.Storage().Get(filename)
		defer readCloser.Close()

		if err != nil {
			log.Println(err)
			return
		}

		if !ok {
			log.Println("Not found")
			return
		}

		io.Copy(w, readCloser)
	})
}

func ApiPostArtifactHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := uniuri.NewLen(32)

		// save file to temp directory
		tmpFilename := container.Config().UploadTmpDir + "/" + filename
		file, err := os.Create(tmpFilename)

		defer os.Remove(tmpFilename)
		defer file.Close()

		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		_, err = io.Copy(file, r.Body)
		defer r.Body.Close()

		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		// create thumbnail
		file.Seek(0, 0)
		img, _, err := image.Decode(file)

		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}

		thumb := resize.Thumbnail(300, 200, img, resize.Bicubic)

		var buff bytes.Buffer
		jpeg.Encode(&buff, thumb, nil)

		// save image to storage
		file.Seek(0, 0)
		savedFileId, err := container.Storage().Save(file)

		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}

		// save thumb to storage
		savedThumbId, err := container.Storage().Save(&buff)

		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}

		//@todo save file ids with information about user/branch/project etc into db

		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(ApiPostArtifactResponse{FileId: savedFileId, ThumbFileId: savedThumbId})
		fmt.Fprint(w, string(j))

		return
	})
}
