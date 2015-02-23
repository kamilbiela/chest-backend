package httphandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/kamilbiela/chest-backend/lib"
	"golang.org/x/oauth2"
)

func LoginGithubHandler(c *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := c.GithubOauth2Config()

		code := c.Auth().RandStr()

		session, _ := c.Session(r)
		session.Values["github_code"] = code
		session.Save(r, w)

		url := config.AuthCodeURL(code, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, 302)
	})
}

func LoginGithubAcceptedHandler(container *lib.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()

		session, _ := container.Session(r)
		code := session.Values["github_code"]

		log.Println(code)

		if vars.Get("state") != code {
			log.Println("Invalid state")
		}

		config := container.GithubOauth2Config()
		tok, err := config.Exchange(oauth2.NoContext, vars.Get("code"))
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		client := github.NewClient(config.Client(oauth2.NoContext, tok))
		can, err := canAccessSite(container, client)

		if err != nil {
			// @todo - print nice html error page with redirect button to login page, or maybe redirect with error message to angular app
			fmt.Fprint(w, "<html><body>")
			fmt.Fprint(w, err)
			fmt.Fprint(w, "<br><a href='/'>Go Back</a>")
			fmt.Fprint(w, "</body></html>")
			return
		}

		if can {
			token, err := container.Auth().GenToken()

			if err != nil {
				log.Fatalln(err)
			}

			// @todo get current domain from request
			url := "http://localhost:3000/#/auth?token=" + token.Val
			http.Redirect(w, r, url, 302)
		}
	})
}

func canAccessSite(container *lib.Container, client *github.Client) (bool, error) {
	orgs, _, err := client.Organizations.List("", nil)

	if err != nil {
		log.Println(err)
		return false, err
	}

	for _, org := range orgs {
		for _, allowedOrg := range container.Config().Github.AllowedOrganizations {
			if org.Login != nil && *org.Login == allowedOrg {
				return true, nil
			}
		}
	}

	return false, nil
}
