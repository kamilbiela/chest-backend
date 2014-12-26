package lib

import (
	"net/http"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/kamilbiela/chest-backend/cache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Container struct {
	config       *Config
	sessionStore sessions.Store
	db           *sql.DB
	cache        cache.Cacher
	auth         *Auth
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Config() *Config {
	if c.config == nil {
		c.config = NewConfig()
	}
	return c.config
}

func (c *Container) GetGithubOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.config.Github.ClientID,
		ClientSecret: c.config.Github.ClientSecret,
		Scopes:       []string{"user:email", "read:org"},
		Endpoint:     github.Endpoint,
	}
}

func (c *Container) SessionStore() sessions.Store {
	//@todo implement memcache/redis backend (strategy)
	if c.sessionStore == nil {
		c.sessionStore = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	}

	return c.sessionStore
}

func (c *Container) Session(r *http.Request) (*sessions.Session, error) {
	store := c.SessionStore()
	return store.Get(r, "chest")
}

func (c *Container) Database() *sql.DB {
	if c.db == nil {
		db, err := sql.Open("mysql", c.Config().MySQL.DSN)

		if err != nil {
			log.Fatalln(err)
		}

		c.db = db
	}

	return c.db
}

func (c *Container) Cache() cache.Cacher {
	if c.cache == nil {
		if c.Config().Cache.Strategy == "memcache" {
			c.cache = cache.NewMemcache(c.Config().Cache.Memcache.Hosts)
		} else {
			log.Fatalln("No cache configured. Allowed values: [memcache]")
		}
	}

	return c.cache
}

func (c *Container) Auth() *Auth {
	if c.auth == nil {
		c.auth = NewAuth(c.Cache())
	}

	return c.auth
}
