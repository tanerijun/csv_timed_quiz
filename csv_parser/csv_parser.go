// Package csv_parser performs csv parsing through the exported function Parse
package csv_parser

import (
	"encoding/csv"
	"io"
	"os"
)

// Function Parse takes the path to a csv file as argument
// and returns the csv as a slice of slice of strings.
func Parse(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	res := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		res = append(res, record)
	}

	return res, nil
}
