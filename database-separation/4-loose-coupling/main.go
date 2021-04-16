package main

import (
	"net/http"
)

func main() {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	createTable(db)
	insertTestData(db)

	bookService := Service{
		store: db,
	}

	http.HandleFunc("/add", bookService.AddBook)
	http.HandleFunc("/all", bookService.GetAllBooks)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func createTable(db *DB) {
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

func insertTestData(db *DB) {
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
