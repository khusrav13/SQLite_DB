package DBS

import (
	"database/sql"
	"fmt"
)

func init()  {
	fmt.Println("hello, I am test")
}

func DbInit(db *sql.DB) (err error) {
	const usersDB = `CREATE TABLE if not exists users (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text not null,
    surname TEXT NOT NULL,
    sex TEXT,
    age INTEGER NOT NULL,
    remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(usersDB)

	if err != nil {
		return err
	}
	return nil
}



func DbInitCurrency(db *sql.DB) (err error) {
	const currency = `CREATE TABLE if not exists currency (
    Id integer PRIMARY KEY AUTOINCREMENT,
    Name text
)`
	_, err = db.Exec(currency)

	if err != nil {
		return err
	}
	return nil
}

func DbAccounts(db *sql.DB) (err error){
	const accounts  = `CREATE TABLE if not exists accounts (
    ID integer PRIMARY KEY AUTOINCREMENT,
    UserId integer references users(id) not null ,
    Number integer,
    Amount integer,
    Currency integer references currency(Id),
    Remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(accounts)

	if err != nil {
		return err
	}
	return nil
}


