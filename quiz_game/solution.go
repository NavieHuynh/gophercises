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

func getQuizData(filename string) ([][]string, error){
	// open file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// read rows in the file
	csvReader := csv.NewReader(f)
	return csvReader.ReadAll()
}

func startQuiz(rows [][]string, result *Result, ch chan string) {
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
	// get user defined timeout
	defaultTime := flag.Int("timeout", 30, "specify timeout in seconds")
	flag.Parse()
	
	rows, err := getQuizData("problems.csv")

	if err != nil{
		log.Fatal(err)
	}

	result := newResult(len(rows))

	fmt.Printf("Press Enter to start the quiz app. You will have %d seconds to complete %d questions\n", *defaultTime, result.total)
	fmt.Scanln(&userInput)

	// start quiz in goroutine
	ch := make(chan string)
	go startQuiz(rows, result, ch)

	select {
	case <-ch:
		fmt.Printf("Quiz complete\n")
	case <- time.After(time.Second * time.Duration(*defaultTime)):
		fmt.Printf("Quiz timed out\n")
	}

	// print results
	fmt.Printf("Number of correct Answers: %d\n", result.numCorrect)
	fmt.Printf("Number of incorrect Answers: %d\n", result.total - result.numCorrect)
	fmt.Printf("Total number of questions: %d\n", result.total)
}
