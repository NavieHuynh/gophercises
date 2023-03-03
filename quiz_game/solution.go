package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)
type Result struct {
	numCorrect int
	numIncorrect int
	total int
}

func newResult(numCorrect int, numIncorrect int, total int) *Result{
	r := Result{
		numCorrect: numCorrect,
		numIncorrect: numIncorrect,
		total: total,
	}
	return &r
}

func quiz(ch chan *Result) {
	var userInput string
	var numCorrect, numIncorrect int
	// open file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		// Read row
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Get User Input
		fmt.Printf("Write answer for %s\n", row[0])
		fmt.Scanln(&userInput)

		// track Results
		if strings.TrimSpace(userInput) == row[1] {
			numCorrect += 1
		} else {
			numIncorrect += 1
		}
	}
	ch <- newResult(numCorrect, numIncorrect, numCorrect+numIncorrect)
}

func main() {
	var userInput string
	defaultTime := flag.Int("timeout", 30, "specify timeout in seconds")
	flag.Parse()

	fmt.Printf("Press Enter to start the quiz app. You will have %d seconds to complete\n", *defaultTime)
	fmt.Scanln(&userInput)

	ch := make(chan *Result)
	go quiz(ch)

	select {
	case result := <-ch:
		fmt.Printf("Number of correct Answers: %d\n", result.numCorrect)
		fmt.Printf("Number of incorrect Answers: %d\n", result.numIncorrect)
		fmt.Printf("Total number of questions: %d\n", result.total)
	case <- time.After(time.Second * time.Duration(*defaultTime)):
		fmt.Printf("Quiz timed out")
	}

}
