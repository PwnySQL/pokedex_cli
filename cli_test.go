package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " tEsT dat,! fuNc ",
			expected: []string{"test", "dat,!", "func"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual and expected do not have same length for input '%s' which produced '%q': expected %d vs %d actual", c.input, actual, len(c.expected), len(actual))
			continue
		}
		for i, w := range actual {
			expectedWord := c.expected[i]
			if (w != expectedWord) {
				t.Errorf("actual and expected are not the same word for index %d: expected '%s' vs '%s' actual", i, expectedWord, w)
			}
		}
	}
}
