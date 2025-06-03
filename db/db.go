package db

import (
	"database/sql"

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
	createUsersTable := "CREATE TABLE IF NOT EXISTS products( id INTEGER PRIMARY KEY AUTOINCREMENT,	product_name TEXT NOT NULL,	price FLOAT NOT NULL, image TEXT NOT NULL, free_shipping TEXT BOOLEAN NOT NULL,	categories TEXT NOT NULL, sold INT, sale BOOLEAN, rating FLOAT, date_arrived TEXT, gender TEXT NOT NULL, color TEXT NOT NULL)"

	_,err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create products table.")
	}

	// createCartTable := `CREATE TABLE IF NOT EXISTS cart(
	// id INTEGER PRIMARY KEY AUTOINCREMENT,
	// user_id INTEGER NOT NULL,
	// product_id INTEGER NOT NULL,
	// quantity INTEGER NOT NULL,

	// )`
	
}