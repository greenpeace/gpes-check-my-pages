package main

import "regexp"

// searchInString reads a string and returns all matches in the string.
func searchInString(total string, expression string) string {
	r, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	firstMatch := r.FindString(total)
	return firstMatch
}
