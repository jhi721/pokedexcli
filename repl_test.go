package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			fmt.Printf("%s %s", word, expectedWord)

			if word != expectedWord {
				t.Errorf("%s is not %s", expectedWord, word)
			}
		}
	}
}
