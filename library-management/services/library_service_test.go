package services

import (
	"fundamentals-of-go/library-management/models"
	"testing"
)

func setupLibrary() (*Library, models.Book, models.Member) {
	lib := NewLibrary()

	// Sample data
	book := models.Book{ID: 1, Title: "The Go Programming Language", Author: "Donovan & Kernighan", Status: models.StatusAvailable}
	member := models.Member{ID: 101, Name: "John Doe"}

	// Pre-populate the library with some data for testing borrow/return
	lib.books[book.ID] = &book
	lib.members[member.ID] = &member

	return lib, book, member
}

// Test #1: Test adding a book and listing it.
func TestAddAndListAvailableBooks(t *testing.T) {
	lib := NewLibrary() // Start with a fresh, empty library
	bookToAdd := models.Book{ID: 2, Title: "Learning Go", Author: "Jon Bodner", Status: models.StatusAvailable}

	err := lib.AddBook(bookToAdd)
	if err != nil {
		t.Fatalf("AddBook failed unexpectedly: %v", err)
	}

	availableBooks := lib.ListAvailableBooks()
	if len(availableBooks) != 1 {
		t.Fatalf("Expected 1 available book, but got %d", len(availableBooks))
	}

	if availableBooks[0].ID != bookToAdd.ID {
		t.Errorf("Expected book ID %d, but got %d", bookToAdd.ID, availableBooks[0].ID)
	}
}

// Test #2: Test adding a book that already exists
func TestAddExistingBook(t *testing.T) {
	lib, book, _ := setupLibrary()

	err := lib.AddBook(book)
	if err == nil {
		t.Fatal("Expected an error when adding a book with a duplicate ID, but got nil")
	}
}