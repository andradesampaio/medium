package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestService_GetAllBooks(t *testing.T) {
	mockStore := &Mock{
		books: []Book{
			{Name: "test", Author: "test"},
		},
	}

	service := Service{
		store: mockStore,
	}

	recorder := httptest.NewRecorder()
	service.GetAllBooks(recorder, httptest.NewRequest("GET", "/", nil))

	var books []Book
	err := json.NewDecoder(recorder.Body).Decode(&books)
	if err != nil {
		t.Error(err)
	}

	if books[0] != mockStore.books[0] {
		t.Fail()
	}
}

func TestService_AddBook(t *testing.T) {
	mockStore := &Mock{
		books: make([]Book, 0),
	}

	service := Service{
		store: mockStore,
	}

	book := Book{
		Name:   "test",
		Author: "test",
	}

	b, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	recorder := httptest.NewRecorder()
	service.AddBook(recorder, httptest.NewRequest("GET", "/", bytes.NewBuffer(b)))

	if len(mockStore.books) == 0 {
		t.Fail()
	}

	if mockStore.books[0] != book {
		t.Fail()
	}
}
