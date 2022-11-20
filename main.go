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

	quizzes, err := csv_parser.Parse(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to csv_timed_quiz!")
	fmt.Println("You have _ seconds to answer each quiz.")
	fmt.Println("Press any key to begin.")

	if *shufflePtr {
		quiz_handler.Shuffle(quizzes)
	}

	score := quiz_handler.Run(quizzes, *timePtr)

	fmt.Println("Congratulations! You finished the quiz.")
	fmt.Printf("Your score is: %d/%d\n", score, len(quizzes))
}
