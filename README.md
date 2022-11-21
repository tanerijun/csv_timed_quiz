# CSV Timed Quiz

A quiz CLI app written in Go.

The quizzes are fetched from a csv file.

And there is also a timer mechanism.

TODO:

- Timer
- Make parser return a slice of struct instead
- Handle taking input from user better `b, err := ioutil.ReadAll(os.Stdin)` or `bufio.NewScanner, scanner.Scan()`
- Also normalize the parsed csv
- Use TrimSpace in fn normalize
