//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name           string
	checkedOut     bool
	lastCheckedOut time.Time
	lastReturned   time.Time
}

func checkOutOrReturnBook(book *Book) {
	now := time.Now().UTC()
	if book.checkedOut {
		book.lastReturned = now
	} else {
		book.lastCheckedOut = now
	}
	book.checkedOut = !book.checkedOut
}

func printLibrary(books []Book) {
	var checkedOut string
	var lastAction string
	for _, book := range books {
		if !book.checkedOut {
			checkedOut = "is not"
			lastAction = "It was returned at " + book.lastReturned.String() + "."
		} else {
			checkedOut = "is"
			lastAction = "It was last checked out at " + book.lastCheckedOut.String()
		}
		fmt.Println("Book", "\""+book.name+"\"", checkedOut, "checked out.", lastAction)
	}
	fmt.Println("There are", len(books), "books")
}

func main() {
	var books []Book
	var now = time.Now().UTC()
	books = append(books, Book{name: "The First Book", checkedOut: false, lastReturned: now, lastCheckedOut: now})
	books = append(books, Book{name: "The Second Book", checkedOut: false})
	books = append(books, Book{name: "The Third Book", checkedOut: false})
	books = append(books, Book{name: "The Fourth Book", checkedOut: false})

	printLibrary(books)

	// Check out a book
	checkOutOrReturnBook(&books[1])
	printLibrary(books)

	//Return the book
	checkOutOrReturnBook(&books[1])
	printLibrary(books)

}
