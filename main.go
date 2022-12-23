package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	date := time.Now().Format("2006-01-02")

	res, err := http.Get("https://www.nytimes.com/svc/wordle/v2/" + date + ".json")
	if err != nil {
		log.Fatal(err)
	}

	s := struct {
		ID              int    `json:"id"`
		Solution        string `json:"solution"`
		PrintDate       string `json:"print_date"`
		DaysSinceLaunch int    `json:"days_since_launch"`
		Editor          string `json:"editor"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Today's wordle is:", s.Solution)
}
