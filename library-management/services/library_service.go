package services

import (
	"fmt"
	"fundamentals-of-go/library-management/models"
)

type Library struct {
	books map[int]*models.Book
	members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[int]*models.Book),
		members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) error {
	if _, exist := l.books[book.ID]; exist {
		return fmt.Errorf("The book with Book ID %d exists", book.ID)
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
