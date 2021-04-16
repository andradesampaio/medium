package main

import (
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare(`select * from books`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results, err := stmt.Query()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer results.Close()

	books := make([]Book, 0)
	err = results.Scan(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
