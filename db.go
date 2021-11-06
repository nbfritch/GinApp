package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

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

type DbContext struct {
	DB      *sql.DB
	Queries map[string]string
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

func nameQuery(file string) string {
	s1 := strings.ReplaceAll(file, "sql/", "")
	s2 := strings.ReplaceAll(s1, ".sql", "")
	return strings.ReplaceAll(s2, "/", ".")
}

func LoadQueriesFromDisk() map[string]string {
	matches, err := filepath.Glob("sql/**/*.sql")
	if err != nil {
		log.Fatal(err)
	} else if len(matches) == 0 {
		log.Fatal(errors.New("Failed to load queries"))
	}

	queries := make(map[string]string)

	for _, m := range matches {
		queryName := nameQuery(m)
		_, hasValue := queries[queryName]
		if hasValue {
			log.Fatal(errors.New(fmt.Sprintf("Attempted to load two queries with the same name: %s", m)))
		}
		content, err := ioutil.ReadFile(m)
		if err != nil {
			log.Fatal(err)
		}
		queries[queryName] = string(content)
		fmt.Printf("Loaded query %s as %s\n", m, queryName)
	}

	return queries
}
