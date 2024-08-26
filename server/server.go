package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Open("file.txt")

	f1, err := os.Open("file2.txt")

	if err != nil {

	}

	log.Fatal(f1, f)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
