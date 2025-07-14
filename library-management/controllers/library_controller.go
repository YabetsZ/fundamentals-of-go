package controllers

import (
	"bufio"
	"fmt"
	"fundamentals-of-go/library-management/models"
	"os"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}


// LibraryController handles the user interaction logic.
type LibraryController struct {
	service LibraryManager
	reader  *bufio.Reader
}

// NewLibraryController is the constructor for our controller.
func NewLibraryController(service LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin), // Reads from standard console input
	}
}

// Run starts the main application loop.
func (c *LibraryController) Run() {
	fmt.Println("Welcome to the Library Management System!")

	// The main application loop
	for {
		c.showMenu()
		choice := c.getUserChoice()

		switch choice {
		case 1:
			c.addBook()
		case 2:
			c.removeBook()
		case 3:
			c.borrowBook()
		case 4:
			c.returnBook()
		case 5:
			c.listAvailableBooks()
		case 6:
			c.listBorrowedBooks()
		case 7:
			fmt.Println("Thank you for using the Library Management System. Goodbye!")
			return // Exit the loop and the function
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println("----------------------------------------")
	}
}
