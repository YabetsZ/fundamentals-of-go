package services

import (
	"library-management/models"
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

// Test #3: Test removing a book
func TestRemoveBook(t *testing.T) {
	lib, book, _ := setupLibrary()

	// Test successful removal
	err := lib.RemoveBook(book.ID)
	if err != nil {
		t.Fatalf("RemoveBook failed unexpectedly: %v", err)
	}

	if _, exists := lib.books[book.ID]; exists {
		t.Error("Book should have been removed, but it still exists in the map.")
	}

	// Test removing a non-existent book
	err = lib.RemoveBook(999) // A non-existent ID
	if err == nil {
		t.Fatal("Expected an error when removing a non-existent book, but got nil")
	}
}

// Test #4: Test borrowing a book
func TestBorrowBook(t *testing.T) {
	t.Run("Successful Borrow", func(t *testing.T) {
		lib, book, member := setupLibrary()

		err := lib.BorrowBook(book.ID, member.ID)
		if err != nil {
			t.Fatalf("BorrowBook failed unexpectedly: %v", err)
		}

		// Check 1: The book's status in the main library list must be "Borrowed".
		if lib.books[book.ID].Status != models.StatusBorrowed {
			t.Errorf("Expected book status to be '%s', but got '%s'", models.StatusBorrowed, lib.books[book.ID].Status)
		}

		// Check 2: The book must be in the member's borrowed list.
		if len(lib.members[member.ID].BorrowedBooks) != 1 {
			t.Fatalf("Expected member's borrowed list to have 1 book, but got %d", len(lib.members[member.ID].BorrowedBooks))
		}
		if lib.members[member.ID].BorrowedBooks[0].ID != book.ID {
			t.Errorf("Member borrowed the wrong book. Expected ID %d, got %d", book.ID, lib.members[member.ID].BorrowedBooks[0].ID)
		}
	})

	t.Run("Borrow a non-existent book", func(t *testing.T) {
		lib, _, member := setupLibrary()
		err := lib.BorrowBook(999, member.ID)
		if err == nil {
			t.Error("Expected an error when borrowing a non-existent book, but got nil")
		}
	})

	t.Run("Borrow with a non-existent member", func(t *testing.T) {
		lib, book, _ := setupLibrary()
		err := lib.BorrowBook(book.ID, 999)
		if err == nil {
			t.Error("Expected an error when borrowing with a non-existent member, but got nil")
		}
	})

	t.Run("Borrow an already borrowed book", func(t *testing.T) {
		lib, book, member := setupLibrary()
		// First borrow is successful
		_ = lib.BorrowBook(book.ID, member.ID)

		// Try to borrow it again
		err := lib.BorrowBook(book.ID, member.ID)
		if err == nil {
			t.Error("Expected an error when borrowing an already borrowed book, but got nil")
		}
	})
}

// Test #5: Test returning a book
func TestReturnBook(t *testing.T) {
	t.Run("Successful Return", func(t *testing.T) {
		lib, book, member := setupLibrary()
		// We must borrow it first to be able to return it
		_ = lib.BorrowBook(book.ID, member.ID)

		err := lib.ReturnBook(book.ID, member.ID)
		if err != nil {
			t.Fatalf("ReturnBook failed unexpectedly: %v", err)
		}

		// Check 1: Book status is now "Available"
		if lib.books[book.ID].Status != models.StatusAvailable {
			t.Errorf("Expected book status to be '%s', but got '%s'", models.StatusAvailable, lib.books[book.ID].Status)
		}

		// Check 2: Member's borrowed list is now empty
		if len(lib.members[member.ID].BorrowedBooks) != 0 {
			t.Errorf("Expected member's borrowed list to be empty, but it has %d books", len(lib.members[member.ID].BorrowedBooks))
		}
	})

	t.Run("Return a book that was not borrowed", func(t *testing.T) {
		lib, book, member := setupLibrary() // Book is "Available" by default
		err := lib.ReturnBook(book.ID, member.ID)
		if err == nil {
			t.Error("Expected an error when returning a book that was not borrowed, but got nil")
		}
	})

	t.Run("Return a book for the wrong member", func(t *testing.T) {
		lib, book, member1 := setupLibrary()

		// Create a second member
		member2 := models.Member{ID: 102, Name: "Jane Smith"}
		lib.members[member2.ID] = &member2

		// Member 1 borrows the book
		_ = lib.BorrowBook(book.ID, member1.ID)

		// Member 2 tries to return it
		err := lib.ReturnBook(book.ID, member2.ID)
		if err == nil {
			t.Error("Expected an error when the wrong member tries to return a book, but got nil")
		}
	})
}

// Test #6: Test listing borrowed books by a member
func TestListBorrowedBooks(t *testing.T) {
	lib, book, member := setupLibrary()
	// Borrow the book to set up the test state
	_ = lib.BorrowBook(book.ID, member.ID)

	borrowedBooks, err := lib.ListBorrowedBooks(member.ID)
	if err != nil {
		t.Fatalf("ListBorrowedBooks failed unexpectedly: %v", err)
	}

	if len(borrowedBooks) != 1 {
		t.Fatalf("Expected to list 1 borrowed book, but got %d", len(borrowedBooks))
	}

	if borrowedBooks[0].ID != book.ID {
		t.Errorf("Listed wrong book. Expected ID %d, got %d", book.ID, borrowedBooks[0].ID)
	}

	// Test listing for a non-existent member
	_, err = lib.ListBorrowedBooks(999)
	if err == nil {
		t.Fatal("Expected an error when listing books for a non-existent member, but got nil")
	}
}
