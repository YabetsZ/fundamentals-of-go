package models

import (
	"testing"
)

func TestMember(t *testing.T) {
	member := Member{
		ID:   101,
		Name: "John Doe",
		// BorrowedBooks slice is implicitly nil here, let's test that
	}

	if member.Name != "John Doe" {
		t.Errorf("Expected member name to be 'John Doe', but got '%s'", member.Name)
	}

	// test: a new member has zero borrowed books.
	if len(member.BorrowedBooks) != 0 {
		t.Errorf("Expected new member to have 0 borrowed books, but got %d", len(member.BorrowedBooks))
	}

	// test: adding a book to the slice
	bookToAdd := Book{ID: 1, Title: "Test Book"}
	member.BorrowedBooks = append(member.BorrowedBooks, bookToAdd)

	if len(member.BorrowedBooks) != 1 {
		t.Errorf("Expected member to have 1 borrowed book after adding, but got %d", len(member.BorrowedBooks))
	}
}
