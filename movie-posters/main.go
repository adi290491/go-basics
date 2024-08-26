package main

import (
	"fmt"
	"log"
	"movie-posters/downloader"
	"movie-posters/omdb"
	"os"
)

func main() {
	result, err := omdb.SearchMovie(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result:\n%v", result)

	fmt.Printf("**********Downloading Poster************\n")
	ok, err := downloader.DownloadPoster(result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok)
}
