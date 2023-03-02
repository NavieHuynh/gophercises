package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func quiz() (int, int, int, error) {
	var userInput string
	var numCorrect, numIncorrect int
	// open file
	f, err := os.Open("problems.csv")
	if err != nil {
		return 0, 0, 0, err
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
			return 0, 0, 0, err
		}

		// Get User Input
		fmt.Printf("Write answer for %s\n", row[0])
		fmt.Scanln(&userInput)

		// track results
		if strings.TrimSpace(userInput) == row[1] {
			numCorrect += 1
		} else {
			numIncorrect += 1
		}
	}
	return numCorrect, numIncorrect, numCorrect + numIncorrect, nil
}

func main() {
	var userInput string
	defaultTime := flag.Int("timeout", 30, "specify timeout in seconds")
	flag.Parse()

	fmt.Printf("Press Enter to start the quiz app. You will have %d seconds to complete\n", *defaultTime)
	fmt.Scanln(&userInput)

	numCorrect, numIncorrect, total, err := quiz()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of correct Answers: %d\n", numCorrect)
	fmt.Printf("Number of incorrect Answers: %d\n", numIncorrect)
	fmt.Printf("Total number of questions: %d\n", total)
}
