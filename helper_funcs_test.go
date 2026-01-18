package main

import ("testing")

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},{
			input: "Hola Mundo  ",
			expected: []string{"hola", "mundo"},
		},{
			input: "THE END IS NEVER THE END IS NEVER THE END",
			expected: []string{"the", "end", "is", "never", "the", "end", "is", "never", "the", "end"},
		},{
			input: " tOtAlLy VaLiD iNpUt",
			expected: []string{"totally", "valid", "input"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %v, got length: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word) 
			}
	}
}}
