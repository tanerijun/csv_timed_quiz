package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tanerijun/csv_timed_quiz/handler"
	"github.com/tanerijun/csv_timed_quiz/parser"
)

func main() {
	// Handle flags
	shufflePtr := flag.Bool("s", false, "Shuffle quizzes")
	timePtr := flag.Int("t", 30, "Time limit for each question")
	flag.Parse()
	filePath := flag.Arg(0)

	parsed, err := parser.Parse(filePath)
	if err != nil {
		exitWithError(err)
	}

	quizzes := handler.NewQuizzesFromSlice(parsed)
	if *shufflePtr {
		handler.Shuffle(quizzes)
	}

	fmt.Println("Welcome to csv_timed_quiz!")
	fmt.Println("You have _ seconds to answer each quiz.")
	fmt.Println("Press any key to begin.")

	score, err := handler.Run(quizzes, *timePtr)
	if err != nil {
		exitWithError(err)
	}

	fmt.Println("Congratulations! You finished the quiz.")
	fmt.Printf("Your score is: %d/%d\n", score, len(quizzes))
}

func exitWithError(e error) {
	fmt.Println("Oops! Something went wrong.")
	fmt.Println("The program is terminated with the following error:")
	log.Fatal(e)
}
