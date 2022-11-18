package main

import (
	"fmt"

	"github.com/tanerijun/csv_timed_quiz/csv_parser"
	"github.com/tanerijun/csv_timed_quiz/quiz_handler"
)

func main() {
	quizzes, err := csv_parser.Parse("quizzes/basic_math.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to csv_timed_quiz!")
	fmt.Println("You have _ seconds to answer each quiz.")
	fmt.Println("Press any key to begin.")

	fmt.Println(quizzes)
	quiz_handler.Shuffle(quizzes)
	fmt.Println(quizzes)

	// score := quiz_handler.Run(quizzes)

	// fmt.Println("Congratulations! You finished the quiz.")
	// fmt.Printf("Your score is: %d/%d\n", score, len(quizzes))
}
