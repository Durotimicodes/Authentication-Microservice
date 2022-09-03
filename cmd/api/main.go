package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"
	_"github.com/jackc/pgconn"
	_"github.com/jackc/pgx/v4"
	_"github.com/jackc/pgx/v4/stdlib"

	"github.com/durotimicodes/authentication/data"
)

const webPort = "83"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	//connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	//set-up config
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

//connect to the database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//function to connect to DB
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			count++
		} else {
			log.Println("Connected to postgres...")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
