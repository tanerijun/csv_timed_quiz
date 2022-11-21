package main

import (
	"flag"
	"fmt"

	"github.com/tanerijun/csv_timed_quiz/csv_parser"
	"github.com/tanerijun/csv_timed_quiz/quiz_handler"
)

func main() {
	// Handle flags
	shufflePtr := flag.Bool("s", false, "Shuffle quizzes")
	timePtr := flag.Int("t", 30, "Time limit for each question")
	flag.Parse()
	filePath := flag.Arg(0)

	parsed, err := csv_parser.Parse(filePath)
	if err != nil {
		panic(err)
	}

	quizzes := quiz_handler.NewQuizzesFromSlice(parsed)
	if *shufflePtr {
		quiz_handler.Shuffle(quizzes)
	}

	fmt.Println("Welcome to csv_timed_quiz!")
	fmt.Println("You have _ seconds to answer each quiz.")
	fmt.Println("Press any key to begin.")

	score, err := quiz_handler.Run(quizzes, *timePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Congratulations! You finished the quiz.")
	fmt.Printf("Your score is: %d/%d\n", score, len(quizzes))
}
