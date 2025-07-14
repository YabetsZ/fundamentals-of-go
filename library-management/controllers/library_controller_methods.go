package controllers

import (
	"fmt"
	"fundamentals-of-go/library-management/models"
	"strconv"
	"strings"
)

// Private Helper Methods

func (c *LibraryController) showMenu() {
	fmt.Println("\nPlease choose an option:")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by a member")
	fmt.Println("7. Exit")
	fmt.Print("Enter your choice: ")
}

// Helper to get integer input from the user
func (c *LibraryController) getIntInput(prompt string) (int, error) {
	fmt.Print(prompt)
	inputStr, _ := c.reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	return strconv.Atoi(inputStr)
}

// Helper to get string input from the user
func (c *LibraryController) getStringInput(prompt string) string {
	fmt.Print(prompt)
	inputStr, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(inputStr)
}

func (c *LibraryController) getUserChoice() int {
	choice, err := c.getIntInput("")
	if err != nil {
		return -1 // Return an invalid choice on error
	}
	return choice
}

func (c *LibraryController) addBook() {
	fmt.Println("\n--- Add a New Book ---")
	id, err := c.getIntInput("Enter Book ID: ")
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	title := c.getStringInput("Enter Title: ")
	author := c.getStringInput("Enter Author: ")

	book := models.Book{ID: id, Title: title, Author: author, Status: models.StatusAvailable}
	err = c.service.AddBook(book)
	if err != nil {
		fmt.Printf("Error adding book: %v\n", err)
	} else {
		fmt.Println("Book added successfully!")
	}
}

func (c *LibraryController) removeBook() {
	fmt.Println("\n--- Remove a Book ---")
	id, err := c.getIntInput("Enter Book ID to remove: ")
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	err = c.service.RemoveBook(id)
	if err != nil {
		fmt.Printf("Error removing book: %v\n", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func (c *LibraryController) borrowBook() {
	fmt.Println("\n--- Borrow a Book ---")
	bookID, err := c.getIntInput("Enter Book ID to borrow: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}
	memberID, err := c.getIntInput("Enter your Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	err = c.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error borrowing book: %v\n", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) returnBook() {
	fmt.Println("\n--- Return a Book ---")
	bookID, err := c.getIntInput("Enter Book ID to return: ")
	if err != nil {
		fmt.Println("Invalid Book ID.")
		return
	}
	memberID, err := c.getIntInput("Enter your Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	err = c.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error returning book: %v\n", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) listAvailableBooks() {
	fmt.Println("\n--- Available Books ---")
	books := c.service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books are currently available.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) listBorrowedBooks() {
	fmt.Println("\n--- List Member's Borrowed Books ---")
	memberID, err := c.getIntInput("Enter your Member ID: ")
	if err != nil {
		fmt.Println("Invalid Member ID.")
		return
	}

	books, err := c.service.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error listing books: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("This member has not borrowed any books.")
		return
	}

	fmt.Printf("Books borrowed by member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}