package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenConnection(url, database, username, password string) {
	connectionDbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", username,password, url, database)

	db, err := sql.Open("mysql", connectionDbStr)

	if err != nil {
		log.Fatalln("Erro ao conectar com o banco")
	}


	if err = db.Ping(); err != nil {
		log.Fatalln("Erro ao conectar com o banco")
	}


	DB = db
}
