package main

import (
	"fmt"
	"issue-tracker/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {

		var age string
		if time.Now().Year()-item.CreatedAt.Year() > 1 {
			age = "more than a year old"
		} else if int(time.Now().Month())-int(item.CreatedAt.Month()) > 1 {
			age = "less than a year old"
		} else {
			age = "less than a month old"
		}

		fmt.Printf("#%-5d %9.9s %.55s %.20s %10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt, age)
	}
}
