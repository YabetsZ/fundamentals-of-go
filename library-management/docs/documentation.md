# Library Management System Documentation

## Overview

This project is a simple, console-based library management system written in Go. It demonstrates key Go features and professional software design principles, including layered architecture, interfaces, and dependency injection. The system allows users to manage a collection of books and track which members have borrowed them.

## Features

-   Add new books to the library collection.
-   Remove books from the collection by their ID.
-   Allow registered members to borrow available books.
-   Allow members to return books they have borrowed.
-   List all books currently available for borrowing.
-   List all books currently borrowed by a specific member.

## Project Structure

The project follows a clean, layered architecture to separate concerns:
```
library_management/
├── main.go # Entry point, composes the application
├── controllers/
│   └── library_controller.go # Handles user input and presentation
├── models/
│   ├── book.go # Defines the Book struct and constants
│   └── member.go # Defines the Member struct
├── services/
│   └── library_service.go # Contains all business logic and data manipulation
├── docs/
│   └── documentation.md # This file
└── go.mod # Go module definition
```
-   **models**: Contains the core data structures (`Book`, `Member`). Has no business logic.
-   **services**: Contains the application's "brain". All business rules (e.g., a book must be "Available" to be borrowed) are enforced here.
-   **controllers**: Acts as the bridge between the user and the service layer. It handles console I/O and calls the appropriate service methods.

## How to Run

### Prerequisites

-   Go (version 1.18 or later) must be installed on your system.

### Steps

1.  **Clone or Download the Repository:**
    Get the project code onto your local machine.

2.  **Navigate to the Project Directory:**
    Open your terminal and `cd` into the `library_management` folder.
    ```bash
    cd /path/to/library_management
    ```

3.  **Run the Tests (Recommended):**
    Ensure everything is working correctly by running the test suite. The `...` wildcard runs tests in all subdirectories.
    ```bash
    go test -v ./...
    ```

4.  **Run the Application:**
    Execute the `main.go` file to start the application.
    ```bash
    go run main.go
    ```