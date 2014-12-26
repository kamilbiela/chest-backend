// GenToken
// CheckToken
// middleware checking token
package lib

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kamilbiela/chest-backend/cache"
	"github.com/kamilbiela/chest-backend/model"
)

type Auth struct {
	cache cache.Cacher
}

func NewAuth(c cache.Cacher) *Auth {
	return &Auth{
		cache: c,
	}
}

func (a *Auth) GenToken() (*model.Token, error) {
	token := model.NewToken()
	err := a.cache.Set("tok_"+token.Val, []byte("1"), int(token.ExpireAt.Unix()))

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return token, nil
}

func (a *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		v := req.URL.Query()
		tokenVal, ok := v["token"]

		if !ok {
			writeAuthError(res, "Not authorized")
			return
		}

		_, err := a.cache.Get("tok_" + tokenVal[0])

		if err != nil {
			writeAuthError(res, "Not authorized")
			return
		}

		next.ServeHTTP(res, req)
	})
}

func (a *Auth) RandStr() string {
	data := make([]byte, 32)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatalln(err)
	}

	return hex.EncodeToString(data)
}

type authErrorResponse struct {
	Message string
}

func writeAuthError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(authErrorResponse{Message: msg})
	fmt.Fprint(w, string(j))
}
