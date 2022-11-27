# CSV Timed Quiz

A quiz CLI app written in Go.

## Explanation

The app fetch quizzes from a csv file.

You can create your own or use the samples inside the quizzes folder.

Note that the quiz doesn't have to be about math. As long as the format is correct, the topic doesn't matter.

You also don't have to worry about capitalization for the answer, as it's not _case sensitive_.

Example:

```
{question},{answer}
What is the shape of the earth?,round
3+3,6
What is the capital of Japan?,Tokyo
```

## Usage

First, make sure that [Go](https://go.dev/) is installed on your local machine. Then build the binary:

```go
go build .
```

Run the binary with `--help` flag for a more detailed instruction.

- `-f` flag is a path to the csv quiz file.
- `-s` flag controls whether the quizzes need to be shuffled or not.
- `-t` flag controls the timer for each question (default to 10s). If the user is out of time, then the answer is considered wrong and no score is given.

Example:

```
./csv_timed_quiz -s -t=5 -f=quizzes/addition.csv
```

Alternatively, you can also run the `main.go` file directly:

```
go run main.go -s -t=5 -f=quizzes/addition.csv
```
