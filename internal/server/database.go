package server

import (
	"github.com/go-pg/pg/v10"
)

func DBOptions() *pg.Options {
	return &pg.Options{
		Addr:     "localhost:5432",
		Database: "cookies",
		User:     "postgres",
		Password: "postgres",
	}
}

var db *pg.DB

// TODO: Add migration file
func SetConnection(dbOpts *pg.Options) {
	// connect
	db = pg.Connect(dbOpts)
}

func GetDBConnection() *pg.DB { return db }
