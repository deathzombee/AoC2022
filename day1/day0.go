package main

import (
	"fmt"
	"log"
	"os"
    "day1/grab"
)

func main() {
	// if input.txt does not exist run grab.go
	if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
		log.Println("input.txt does not exist, running grab.go")
		grab.Grab()

	}
	//open input.txt and read it into a string
	   content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))



}