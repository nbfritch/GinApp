package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type connectionString struct {
	database string
	User     string
	Host     string
	Password string
	Port     int
	DbName   string
}

func SetupDb() *sql.DB {
	connFile, err := ioutil.ReadFile("./db.json")
	if err != nil {
		log.Fatal(err)
	}

	var c connectionString
	if err := json.Unmarshal(connFile, &c); err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
