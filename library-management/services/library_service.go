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