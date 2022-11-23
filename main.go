package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tanerijun/csv_timed_quiz/handler"
	"github.com/tanerijun/csv_timed_quiz/parser"
)

func main() {
	// Handle flags
	shufflePtr := flag.Bool("s", false, "Whether to shuffle the quizzes or not. (default true)")
	timePtr := flag.Int("t", 10, "Time limit for each question.")
	filePathPtr := flag.String("f", "", "A file path to a csv file with \"question,answer\" format.")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Invalid arguments provided.")
		fmt.Println("Use the help flag (-h or --help) for instructions on how to use the program.")
		os.Exit(1)
	}

	parsed, err := parser.Parse(*filePathPtr)
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
