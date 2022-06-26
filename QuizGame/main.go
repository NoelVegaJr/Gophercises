package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	quizFile := flag.String("quiz", "problems.csv", "File path to quiz. Must be in csv format with no headers and two fields. Field #1 is the question Field #2 is the answer.")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds.")
	flag.Parse()

	correct := 0
	quiz := getQuiz(quizFile)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range quiz {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s = ", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			goto End
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
End:
	fmt.Printf("You scored %d out of %d", correct, len(quiz))
}

type problem struct {
	q string
	a string
}

func parseProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{q: line[0], a: line[1]}
	}
	return problems
}

func getQuiz(path *string) []problem {
	file, _ := os.Open(*path)
	lines, err := csv.NewReader(file).ReadAll()
	if err != err {
		log.Fatal(err)
	}

	return parseProblems(lines)

}
