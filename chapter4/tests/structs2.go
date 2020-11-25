package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title           string `json:"Title,omitempty"`  // only necessary because of omitempty
	Author          string `json:"author,omitempty"` // necessary because the field name and tag are different
	PublicationDate int    `json:"publication_date,omitempty"`
	Owned           bool   `json:"do_i_have_it,omitempty"`
}

func booksJSON() {

	var books = []Book{
		{
			Title:           "Stranger In A Strange Land",
			Author:          "Robert A. Heinlein",
			PublicationDate: 1961,
			Owned:           true,
		},
		{
			Title:           "1984",
			Author:          "George Orwell",
			PublicationDate: 1949,
			Owned:           true,
		},
		{
			Title:           "Brave New World",
			Author:          "Aldous Huxley",
			PublicationDate: 1932,
			Owned:           false,
		},
	}

	b, err := json.Marshal(books)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	fmt.Printf("%v\n\n", b) // prints a slice of bytes
	fmt.Printf("%s\n\n", b) // equivalent to fmt.Println(string(b))

	bIndented, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	fmt.Printf("%v\n\n", bIndented) // prints a slice of bytes
	fmt.Printf("%s\n\n", bIndented) // equivalent to fmt.Println(string(bIndented))

	var myTinyLibrary []Book
	if err := json.Unmarshal(b, &myTinyLibrary); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	fmt.Printf("%v\n", myTinyLibrary)

	var myTinyLibraryFromIndented []Book
	if err := json.Unmarshal(bIndented, &myTinyLibraryFromIndented); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	fmt.Printf("%v\n", myTinyLibraryFromIndented)

}
