package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// only exported struct fields will be marshalled
type Film struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {

	var films = []Film{
		{
			Title: "2001: A Space Odyssey", Year: 1968, Color: true,
			Actors: []string{"Keir Dullea"},
		},
		{
			Title: "A Streetcar Named Desire", Year: 1951, Color: false,
			Actors: []string{"Marlon Brando", "Vivien Leigh"},
		},
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title: "Ladri Di Biciclette", Year: 1948, Color: false,
			Actors: []string{"Enzo Staiola", "Lamberto Maggiorani"},
		},
	}

	// json.Marshal() converts a data structure to JSON and produces a []byte and an error
	d, err := json.Marshal(films)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	fmt.Printf("%s\n\n", d)

	// json.MarshalIndent() produces neatly indented output
	dNew, err := json.MarshalIndent(films, "", "	")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	fmt.Printf("%s\n", dNew)

	var titles []struct{ Title string }
	if err := json.Unmarshal(d, &titles); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	fmt.Println(titles[:])

}
