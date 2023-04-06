package db

import (
	"clean/pkg/config"
	"database/sql"
	"log"
)

func Connectdb(c config.Config) *sql.DB {
	dbname := c.DBName
	dburi := c.DBSOURCE
	db, err := sql.Open("postgres", dburi)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error in ping", err)
	}
	log.Println("\n Connected to database:", dbname)
	return db
}
