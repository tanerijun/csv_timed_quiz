// Package quiz_handler offers functions to work with quizzes
package handler

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

// Function Run runs a quiz game based on the received parameter and returns a score.
func Run(quizzes []Quiz, t int) (int, error) {
	score := 0
	for _, quiz := range quizzes {
		fmt.Print(quiz.question, " = ")

		var in string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in = scanner.Text()
			break
		}
		if err := scanner.Err(); err != nil {
			return score, err
		}

		ans := normalize(quiz.answer)
		in = normalize(in)

		if in == ans {
			score++
			fmt.Println("Nice!")
		} else {
			fmt.Println("Oops!")
		}
	}

	return score, nil
}

// Function normalize trims whitespaces and lowercase a string.
func normalize(s string) string {
	res := strings.TrimSpace(s)
	res = strings.ToLower(res)
	return res
}

// Function Shuffle shuffles quizzes.
func Shuffle(s []Quiz) {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// Function New returns a new Quiz.
func New(s []string) Quiz {
	return Quiz{s[0], s[1]}
}

// Function NewQuizzesFromSlice takes in a slice of string, and returns a slice of Quiz.
func NewQuizzesFromSlice(s [][]string) []Quiz {
	quizzes := make([]Quiz, len(s))
	for i, v := range s {
		quizzes[i] = New(v)
	}

	return quizzes
}
