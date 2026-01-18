package main

import "strings"

func cleanInput(text string) []string {
	var words []string
	currentWord := ""
	for _, c := range text {
		if c == ' ' {
			if len(currentWord) > 0 {
				words = append(words, strings.ToLower(currentWord))
			} 
			currentWord = ""
		} else {
			currentWord += string(c)
		}
	}
	if len(currentWord) > 0 {
		words = append(words, strings.ToLower(currentWord))
	}
	return words
}

