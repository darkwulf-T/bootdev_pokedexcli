package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input: "This will be a pokedex",
			expected: []string{"this", "will", "be", "a", "pokedex"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("acutal slice is not matching expected slice")
		}
		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s is not %s", word, expectedWord)
			}
		}
	}
}