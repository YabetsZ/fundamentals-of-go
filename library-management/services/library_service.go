package services

import (
	"fmt"
	"library-management/models"
)

type Library struct {
	books   map[int]*models.Book
	members map[int]*models.Member
}

func NewLibrary() *Library {
	startingMember := make(map[int]*models.Member)
	startingMember[1] = &models.Member{ID: 1, Name: "John", BorrowedBooks: make([]models.Book, 0)}
	return &Library{
		books:   make(map[int]*models.Book),
		members: startingMember,
	}
}

func (l *Library) AddBook(book models.Book) error {
	if _, exist := l.books[book.ID]; exist {
		return fmt.Errorf("the book with book id %d exists", book.ID)
	}
	l.books[book.ID] = &book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	bookList := make([]models.Book, 0, len(l.books))
	for _, book := range l.books {
		if book.Status == models.StatusAvailable {
			bookList = append(bookList, *book)
		}
	}
	return bookList
}

func (l *Library) RemoveBook(bookID int) error {
	_, exist := l.books[bookID]
	if !exist {
		return fmt.Errorf("a book with book id %d does not exist", bookID)
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	if _, exist := l.books[bookID]; !exist {
		return fmt.Errorf("the book with id %v does not exist", bookID)
	} else if l.books[bookID].Status == models.StatusBorrowed {
		return fmt.Errorf("the book with id %v is borrowd", bookID)
	} else if _, exist := l.members[memberID]; !exist {
		return fmt.Errorf("the member with id %v does not exist", memberID)
	}

	l.books[bookID].Status = models.StatusBorrowed
	l.members[memberID].BorrowedBooks = append(l.members[memberID].BorrowedBooks, *l.books[bookID])

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	// check if the book exists
	bookFromLib, exist := l.books[bookID]
	if !exist {
		return fmt.Errorf("the book with id %v does not exist", bookID)
	}
	// check if member exist
	member, exist := l.members[memberID]
	if !exist {
		return fmt.Errorf("the book with id %v does not exist", bookID)
	}
	// check if the book is borrowed
	if l.books[bookID].Status == models.StatusAvailable {
		return fmt.Errorf("the book with id %v has not been borrowed", bookID)
	}
	// check if the book is borrowd by the member
	hasBorrowed := false
	for idx, book := range member.BorrowedBooks {
		if book.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[(idx+1):]...)
			bookFromLib.Status = models.StatusAvailable
			hasBorrowed = true
		}
	}
	if !hasBorrowed {
		return fmt.Errorf("the member with id %d has not borrowed the book %d", memberID, bookID)
	}
	return nil
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exist := l.members[memberID]
	if !exist {
		return []models.Book{}, fmt.Errorf("the member with id %v does not exist", memberID)
	}
	return member.BorrowedBooks, nil
}
