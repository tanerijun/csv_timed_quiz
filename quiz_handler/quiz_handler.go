// Package quiz_handler offers functions to work with quizzes
package quiz_handler

import "fmt"

// Function Run runs a quiz game based on the received parameter and returns a score.
func Run(quizzes [][]string) int {
	score := 0
	for _, quiz := range quizzes {
		fmt.Print(quiz[0], " = ")
		var ans string
		fmt.Scanln(&ans)

		if ans == quiz[1] {
			score++
			fmt.Println("Nice!")
		} else {
			fmt.Println("Oops!")
		}
	}

	return score
}
