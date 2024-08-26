package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type ActorTitles struct {
	Title  string
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Avengers - Endgame",
			Year:   2019,
			Color:  true,
			Actors: []string{"Robert Downey Jr", "Chris Evans", "Scarlett Johanssen", "Josh Brolin"},
		},
		{
			Title:  "Deadpool and Wolverine",
			Year:   2024,
			Color:  true,
			Actors: []string{"Ryan Reynolds", "Hugh Jackman"},
		},
		{
			Title:  "Shawshank Redemption",
			Year:   1995,
			Color:  false,
			Actors: []string{"Morgan Freeman"},
		},
	}

	data, err := json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	var titlesAndActors []struct {
		Title  string
		Actors []string
	}

	if err := json.Unmarshal(data, &titlesAndActors); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Printf("%s\n", titlesAndActors)

	var listOfActors []struct{ Actors []string }

	decodedMovies, err := os.ReadFile("movies.json")

	reader := bytes.NewReader(decodedMovies)

	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	if err := json.NewDecoder(reader).Decode(&listOfActors); err != nil {
		log.Fatalf("JSON Decoding failed: %s", err)
	}

	fmt.Println(listOfActors)
}
