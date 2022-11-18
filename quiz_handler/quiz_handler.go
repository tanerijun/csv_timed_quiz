// Package quiz_handler offers functions to work with quizzes
package quiz_handler

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Function Run runs a quiz game based on the received parameter and returns a score.
func Run(quizzes [][]string) int {
	score := 0
	for _, quiz := range quizzes {
		fmt.Print(quiz[0], " = ")
		var in string
		fmt.Scan(&in)

		ans := normalize(quiz[1])
		in = normalize(in)

		if in == ans {
			score++
			fmt.Println("Nice!")
		} else {
			fmt.Println("Oops!")
		}
	}

	return score
}

// Function normalize trims whitespaces and lowercase a string.
func normalize(s string) string {
	res := strings.Trim(s, "\t\n ")
	res = strings.ToLower(res)
	return res
}

// Function Shuffle shuffles quizzes
func Shuffle(s [][]string) {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
