package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	err := os.Remove("./reviews.db")
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("sqlite3", "./reviews.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	createTable()
	insertTestData()

	http.HandleFunc("/all", GetAllBooks)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func createTable() {
	stmt, err := db.Prepare(`
		create table books (
			id integer primary key autoincrement,
			name text not null,
			author text not null
		);
		delete from books;
	`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

}

func insertTestData() {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(`
		insert into books(name, author) values
			('Way of Kings', 'Brandon Sanderson'),
			('The Gunslinger', 'Stephen King'),
			('Leviathan Wakes', 'James S. Corey'),
			('East of Eden', 'John Steinbeck');
	`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	tx.Commit()
}
