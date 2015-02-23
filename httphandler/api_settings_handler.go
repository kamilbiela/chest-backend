package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/kamilbiela/chest-backend/lib"
)

type SettingsResponse struct {
	ApiToken string
}

func ApiGetSettingsHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := container.Config()
		s := SettingsResponse{
			ApiToken: config.Secret,
		}

		data, _ := json.Marshal(s)

		w.Write(data)
	})
}
