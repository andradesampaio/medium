package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Store interface {
	GetAllBooks(context.Context) ([]Book, error)
	AddBook(context.Context, Book) error
}

type Service struct {
	store Store
}

func (s *Service) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := s.store.GetAllBooks(r.Context())
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

func (s *Service) AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.store.AddBook(r.Context(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
