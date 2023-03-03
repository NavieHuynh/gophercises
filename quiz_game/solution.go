package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)
type Result struct {
	numCorrect int
	total int
}

func newResult(total int) *Result{
	r := Result{
		numCorrect: 0,
		total: total,
	}
	return &r
}

func quiz(rows [][]string, result *Result, ch chan string) {
	var userInput string

	for _, row := range rows {
		// Get User Input
		fmt.Printf("Write answer for %s\n", row[0])
		fmt.Scanln(&userInput)

		// track Results
		if strings.TrimSpace(userInput) == row[1] {
			result.numCorrect += 1
		}
	}
	ch <- "ok"
}

func main() {
	var userInput string
	defaultTime := flag.Int("timeout", 30, "specify timeout in seconds")
	flag.Parse()
	// open file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()

	if err != nil{
		log.Fatal(err)
	}

	result := newResult(len(rows))

	fmt.Printf("Press Enter to start the quiz app. You will have %d seconds to complete %d questions\n", *defaultTime, result.total)
	fmt.Scanln(&userInput)

	ch := make(chan string)
	go quiz(rows, result, ch)

	select {
	case <-ch:
		fmt.Printf("Quiz complete\n")
	case <- time.After(time.Second * time.Duration(*defaultTime)):
		fmt.Printf("Quiz timed out\n")
	}

	fmt.Printf("Number of correct Answers: %d\n", result.numCorrect)
	fmt.Printf("Number of incorrect Answers: %d\n", result.total - result.numCorrect)
	fmt.Printf("Total number of questions: %d\n", result.total)
}
