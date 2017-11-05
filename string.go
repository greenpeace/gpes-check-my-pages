package main

import (
	"regexp"
	"time"
)

// searchInString reads a string and returns all matches in the string.
func searchInString(total string, expression string) string {
	r, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	firstMatch := r.FindString(total)
	return firstMatch
}

// nowDateTimeString Returns
func nowDateTimeString() string {
	t := time.Now()
	return t.Format("2006-01-02-15h04m05")
}
