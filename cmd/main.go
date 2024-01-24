package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//Get the filename from the flag -filename
	var fileName *string = flag.String("filename", "questions", "The CSV filename.")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("File not found, run the program with a flag -filename=<< >> ")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	allLines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read the CSV file content")
	}

	// Loop through the CSV file containing question and answer. For each line containing
	// question and answer, take the question and prompt the user to enter the answer.
	// Take the answer from the user and compare with the answer given in the CSV file. Let
	// the user knows if their answer is correct or wrong. Finally give the user the total
	// score.

	var totalCorrect int = 0
	for i, line := range allLines {
		fmt.Printf("\nQuestion %d: %s", i+1, line[0])
		fmt.Println("\nPlease enter your answer: ")
		var input string
		fmt.Scanln(&input)

		if input == line[1] {
			fmt.Printf("Answer %s is Correct!\n", input)
			totalCorrect++
		} else {
			fmt.Printf("Answer %s is Wrong!\n", input)
		}
	}
	fmt.Printf("\n###   Your total score is %d/%d    ###", totalCorrect, len(allLines))
}
