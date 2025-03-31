package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    " hello world ",
		expected: []string{"hello", "world"},
	}}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The result slice has a different size as expected")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected this word: %s but got this one in returned: %s ", expectedWord, word)
			}
		}
	}
}
