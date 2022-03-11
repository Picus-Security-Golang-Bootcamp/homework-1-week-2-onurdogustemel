package main

import (
	"fmt"
	"os"
	"strings"
)

type Books struct {
	arr []string
}

var b = Books{
	arr: []string{"The Giver", "Bleak House", "The Lord of The Rings: Fellowship Of The Ring", "White Fang", "Lord of The Flies", "Wiseguy"},
}

// arr slice contains the name of the books.

//Search function prints the name of the book given as input to the screen.
func search(bookName []string) string {

	for _, value := range b.arr {

		for i := range bookName {
			if strings.ToLower(strings.Join(bookName[i:], " ")) == strings.ToLower(value) {
				return "Book found: " + value
			}
		}
	}
	return "No book found under the name you entered."
}

//list function prints the book list
func list() {

	fmt.Println("\n-----Booklist-----")

	for i, books := range b.arr {
		fmt.Printf("%v- %v\n", i+1, books)
	}
}

func main() {

	//args variable returns the command line arguments as a slice from terminal.
	args := os.Args[1:]

	for i := range args {

		if len(args) >= 2 && args[i] == "search" {

			fmt.Println(search(os.Args[2:]))
			break

		} else if len(args) == 1 && args[i] == "list" {

			list()
			break

		} else {

			fmt.Println("Wrong input.")
			break
		}
	}
}
