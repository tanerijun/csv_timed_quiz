package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expectedRes := [][]string{
		{"5+5", "10"},
		{"7+3", "10"},
		{"1+1", "2"},
		{"8+3", "11"},
		{"1+2", "3"},
	}

	res, err := Parse("testdata/sample.csv")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if !(reflect.DeepEqual(res, expectedRes)) {
		t.Errorf("Expected %v, got %v", expectedRes, res)
	}

	_, err = Parse("testdata/doesnt-exist")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
