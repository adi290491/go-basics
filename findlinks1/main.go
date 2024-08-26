package main

import (
	"fmt"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}

}
