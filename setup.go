package main

import (
	"io/ioutil"
	"log"

	"github.com/kamilbiela/chest-backend/lib"
)

func setup(container *lib.Container) {
	db := container.Database()

	sql, err := ioutil.ReadFile("data/schema.sql")
	if err != nil {
		log.Println("There was problem loading setup sql file: ")
		log.Fatalln(err)
	}

	// @todo import sql statement by statement

	res, err := db.Query(string(sql))

	if err != nil {
		log.Println("There was problem importing setup sql file to database: ")
		log.Fatalln(err)
	}

	res.Close()
}
