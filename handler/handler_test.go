package handler

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	strs := [][]string{
		{"5+5", "10"},
		{"7+3", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
		{"1+2", "3"},
	}

	quizzes := make([]Quiz, len(strs))
	for i, v := range strs {
		quizzes[i] = New(v)
	}

	for i := range quizzes {
		if (quizzes[i].question != strs[i][0]) && (quizzes[i].answer != strs[i][1]) {
			t.Errorf("Expected: %v, got: %v", strs[i], quizzes[i])
		}
	}
}

func TestNewQuizzesFromSlice(t *testing.T) {
	strs := [][]string{
		{"5+5", "10"},
		{"7+3", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
		{"1+2", "3"},
	}

	expectedQuizzes := make([]Quiz, len(strs))
	for i, v := range strs {
		expectedQuizzes[i] = New(v)
	}

	quizzes := NewQuizzesFromSlice(strs)

	if !(reflect.DeepEqual(quizzes, expectedQuizzes)) {
		t.Errorf("Expected: %v, got: %v", expectedQuizzes, quizzes)
	}
}

func TestShuffle(t *testing.T) {
	expectedQuizzes := []Quiz{
		{"5+5", "10"},
		{"7+3", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
		{"1+2", "3"},
	}

	quizzes := make([]Quiz, len(expectedQuizzes))
	copy(quizzes, expectedQuizzes)
	Shuffle(quizzes)

	if reflect.DeepEqual(quizzes, expectedQuizzes) {
		t.Error("Equal value after shuffling")
	}
}

func TestRun(t *testing.T) {
	// Replace stdin and restore it later
	realStdin := os.Stdin
	defer func() {
		os.Stdin = realStdin
	}()

	quizzes := []Quiz{
		{"5+5", "10"},
	}

	// Test correct answer
	tmpFile1 := createTempFile([]byte("10"))
	defer os.Remove(tmpFile1.Name())
	defer tmpFile1.Close()
	os.Stdin = tmpFile1
	score, err := Run(quizzes, 2)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != 1 {
		t.Errorf("Expected score: 1, got: %d", score)
	}

	// Test wrong answer
	tmpFile2 := createTempFile([]byte("11"))
	defer os.Remove(tmpFile2.Name())
	defer tmpFile2.Close()
	os.Stdin = tmpFile2
	score, err = Run(quizzes, 2)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != 0 {
		t.Errorf("Expected score: 0, got: %d", score)
	}

	// Test timeout
	tmpFile3 := createTempFile([]byte("10"))
	defer os.Remove(tmpFile3.Name())
	defer tmpFile3.Close()
	os.Stdin = tmpFile3
	score, err = Run(quizzes, 0)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != 0 {
		t.Errorf("Expected score: 0, got: %d", score)
	}
}

func createTempFile(content []byte) *os.File {
	tmpFile, err := os.CreateTemp("", "tmp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpFile
}
