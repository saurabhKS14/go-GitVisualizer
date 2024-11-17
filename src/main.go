package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

func main() {
	_, err := git.PlainOpen(".")

	if err != nil {
		log.Fatalf("This is not a github repository")
	}

	fmt.Println("Done")
}
