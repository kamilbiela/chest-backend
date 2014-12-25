package main

import (
	"github.com/turbogopher/chest-backend/lib"
)

type Container struct {
	config *lib.Config
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Config() *lib.Config {
	if c.config == nil {
		c.config = lib.NewConfig()
	}
	return c.config
}
