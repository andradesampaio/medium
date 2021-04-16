package main

import "context"

type Mock struct {
	books []Book
}

func (m *Mock) GetAllBooks(ctx context.Context) ([]Book, error) {
	return m.books, nil
}

func (m *Mock) AddBook(ctx context.Context, book Book) error {
	m.books = append(m.books, book)
	return nil
}
