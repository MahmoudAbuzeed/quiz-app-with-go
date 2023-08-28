package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to epen the CSV file: %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	correct := 0
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {

	result := make([]problem, len(lines))
	for index, line := range lines {
		result[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return result

}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
