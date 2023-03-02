package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var userAnswer string
	var numCorrect, numIncorrect int

	fmt.Printf("Starting Quiz App\n")
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
		fmt.Scanln(&userAnswer)

		// track results
		if userAnswer == row[1] {
			numCorrect += 1
		} else {
			numIncorrect += 1
		}
	}

	fmt.Printf("Number of correct Answers: %d\n", numCorrect)
	fmt.Printf("Number of incorrect Answers: %d\n", numIncorrect)
	fmt.Printf("Total number of questions: %d\n", numCorrect+numIncorrect)
}
