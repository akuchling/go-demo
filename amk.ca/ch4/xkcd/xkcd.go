package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type XKCDCartoon struct {
	Index      int
	Title      string
	Transcript string
	Alt, Url   string
}

func main() {
	var cartoons []XKCDCartoon
	var err error

	if len(os.Args) == 1 {
		fmt.Printf("Usage:\n  %s <search terms>\n", os.Args[0])
	}

	file, err := os.Open("xkcd.json")
	if err != nil {
		log.Fatalf("Error reading list: %s", err)
	}

	err = json.NewDecoder(file).Decode(&cartoons)
	for _, cartoon := range cartoons {
		terms_found := 0
		corpus := strings.ToLower(cartoon.Title + " " + cartoon.Alt + " " +
			cartoon.Transcript)
		for j := 1; j < len(os.Args); j++ {
			if strings.Contains(corpus, strings.ToLower(os.Args[j])) {
				terms_found++
			}
		}
		if terms_found == len(os.Args)-1 {
			fmt.Printf("%s %s\n", cartoon.Url, cartoon.Title)
		}
	}
}
