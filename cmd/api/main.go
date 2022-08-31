package main

import (
	"database/sql"
	"log"

	"github.com/durotimicodes/authentication/data"
)

const webPort = "83"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	

}
