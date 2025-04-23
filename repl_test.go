package main

import "testing"
		


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:  "One two   three  fOUR",
			expected: []string{"one", "two", "three", "four"},
		},

		{
			input: "   monKey doLPHIN DRAGON giraffe     lizzard",
			expected: []string{"monkey", "dolphin", "dragon", "giraffe", "lizzard"},
		},
	}
	
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length does not match")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words dont match")
				return
			}
		}
	}


}
