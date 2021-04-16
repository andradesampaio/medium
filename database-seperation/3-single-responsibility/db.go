package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	err := os.Remove("./reviews.db")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", "./reviews.db")
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) GetAllBooks(ctx context.Context) ([]Book, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`select * from books`)
	if err != nil {
		return nil, err
	}

	results, err := stmt.Query()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer results.Close()

	books := make([]Book, 0)
	err = results.Scan(&books)
	if err != nil {
		return nil, err
	}

	return books, err
}

func (db *DB) AddBook(ctx context.Context, book Book) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`insert into books(name, author) values(?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.Name, book.Author)
	return err
}
