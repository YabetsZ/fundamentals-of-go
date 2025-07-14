package models

import (
	"testing"
)

func TestBookStruct(t *testing.T) {
	book := Book{
		ID: 1,
		Title: "Rust for rustaceans",
		Author: "Jon Gjengset",
		Status: StatusAvailable,
	}

	if book.Status != StatusAvailable {
		t.Errorf("Expected book status to be '%s', but got '%s'", StatusAvailable, book.Status)
	}

	if book.Title == "" {
		t.Errorf("Expected book to have a title, but it was empty")
	}
}