package main

import (
	"WorkData/DBS"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	database, err := sql.Open("sqlite3", "BankDB")
	if err != nil{
		log.Fatal("DB isn't work", err)
	}
	_ = DBS.DbInit(database)
	_ = DBS.DbInitCurrency(database)
	_ = DBS.DbAccounts(database)

	if err != nil {
		log.Fatal("error ", err)
	}
	fmt.Println("ok")
}