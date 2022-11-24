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

type Input struct {
	input string
	err   error
}

// Function Run runs a quiz game based on the received parameter and returns a score.
func Run(quizzes []Quiz, t int) (int, error) {
	score := 0
	inCh := make(chan Input)

	for _, quiz := range quizzes {
		fmt.Print(quiz.question, " = ")

		timer := time.NewTimer(time.Duration(t) * time.Second)
		go getInput(inCh)

		select {
		case in := <-inCh:
			if in.err != nil {
				return score, in.err
			}

			usrInput := in.input
			ans := normalize(quiz.answer)
			usrInput = normalize(usrInput)

			if usrInput == ans {
				score++
				fmt.Println("Nice! You got it right.")
			} else {
				fmt.Println("Oops! Your answer is wrong.")
			}
		case <-timer.C:
			fmt.Println("\nOops! You're out of time.")
		}
	}

	return score, nil
}

func getInput(c chan Input) {
	var in string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in = scanner.Text()
	c <- Input{in, scanner.Err()}
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
