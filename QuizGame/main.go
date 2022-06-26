package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func getQuiz(filePath *string) [][]string {
	file, _ := os.Open(*filePath)

	records, err := csv.NewReader(file).ReadAll()

	if err != err {
		log.Fatal(err)

	}

	return records

}

var quizFile = flag.String("quiz", "problems.csv", "File path to quiz. Must be in csv format with no headers and two fields. Field #1 is the question Field #2 is the answer.")

func main() {

	flag.Parse()
	correct := 0

	problems := getQuiz(quizFile)
	for i, problem := range problems {
		var userAnswer string
		question := problem[0]
		correctAnswer := problem[1]

		fmt.Printf("Problem #%d: %s = ", i+1, question)
		fmt.Scanf("%s = ", &userAnswer)

		if userAnswer == correctAnswer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
}
