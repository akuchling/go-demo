package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type XKCDCartoon struct {
	Index      int
	Title      string
	Transcript string
	Alt, Url   string
}

const TOTAL = 1940 // Most recent cartoon as of January 11 2018

func main() {
	var cartoons []XKCDCartoon

	for i := 1; i <= TOTAL; i++ {
		url := "https://www.xkcd.com/" + strconv.Itoa(i) + "/"
		resp, err := http.Get(url + "info.0.json")
		if err != nil {
			log.Fatalf("Error retrieving cartoon #%d: %s", i, err)
		}
		if resp.StatusCode != 200 {
			log.Printf("HTTP status %d for cartoon #%d; skipping",
				resp.StatusCode, i)
			continue
		}

		// Unmarshal data
		cartoon := XKCDCartoon{Url: url}

		err = json.NewDecoder(resp.Body).Decode(&cartoon)
		resp.Body.Close()
		if err != nil {
			log.Fatalf("Error unmarshalling cartoon #%d: %s", i, err)
		}
		cartoons = append(cartoons, cartoon)
	}

	data, err := json.MarshalIndent(cartoons, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling cartoon list: %s", err)
	}
	fmt.Printf("%s", data)
}
