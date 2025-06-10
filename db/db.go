package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "plashoes.db")

	if err != nil {
		panic("Could not create database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables(){
	createProductsTable := "CREATE TABLE IF NOT EXISTS products( id INTEGER PRIMARY KEY AUTOINCREMENT,	product_name TEXT NOT NULL,	price FLOAT NOT NULL, image TEXT NOT NULL, free_shipping TEXT BOOLEAN NOT NULL,	categories TEXT NOT NULL, sold INT, sale BOOLEAN, rating FLOAT, date_arrived TEXT, gender TEXT NOT NULL, color TEXT NOT NULL)"

	_,err := DB.Exec(createProductsTable)

	if err != nil {
		panic("Could not create products table.")
	}

	createUsersTable := "CREATE TABLE IF NOT EXISTS users ( id INTEGER PRIMARY KEY AUTOINCREMENT, user_name TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, mobile_no INT NOT NULL, date_of_birth date NOT NULL, gender TEXT NOT NULL, country TEXT NOT NULL, postal_code INT NOT NULL, address TEXT NOT NULL, country_code TEXT NOT NULL )"

	_,err = DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createCartTable := "CREATE TABLE IF NOT EXISTS cart(id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER NOT NULL,	product_id INTEGER NOT NULL, quantity INTEGER NOT NULL, FOREIGN KEY(user_id) REFERENCES users(id), FOREIGN KEY(product_id) REFERENCES products(id))"
	
	_,err = DB.Exec(createCartTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create cart table.")
	}
}