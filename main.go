package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/hurb-currency/config"
)

func main() {
	fmt.Println("Teste")

	database := "currencies"
	url := "localhost:35060"
	username := "root"
	password := "root"

	config.OpenConnection(url, database, username, password)

	defer config.DB.Close()

	log.Fatal(http.ListenAndServe(":8990", nil))
}
