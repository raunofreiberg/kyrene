package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//"time"

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	conf[dbhost] = dbhost
	conf[dbport] = dbport
	conf[dbuser] = dbuser
	conf[dbpass] = dbpass
	conf[dbname] = dbname

	for key, _ := range conf {
		val, exists := os.LookupEnv(key)
		if !exists {
			panic("oh no")
		}
		conf[key] = val
	}
	return conf
}

func InitDb() {
	config := dbConfig()
	var err error

	psqlInfo := fmt.Sprintf(
		`
		host=%s
		port=%s
		user=%s
		password=%s
		dbname=%s
		sslmode=disable
		`,
		config[dbhost],
		config[dbport],
		config[dbuser],
		config[dbpass],
		config[dbname],
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB =>", config[dbname])

	db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, content TEXT, is_completed BOOL)")
	//db.Exec(`INSERT INTO todos (content, is_completed) VALUES ('do shit', true)`)
}