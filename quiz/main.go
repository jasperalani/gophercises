package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Solution

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question, answer'.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read provided csv file"))
	}

	problems := parseLines(lines)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })

	var correct = 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == p.answer {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect :(")
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
