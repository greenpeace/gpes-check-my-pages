package main

import (
	"encoding/csv"
	"os"
)

// readCsvFile reads a file and returns a [][]string slice
func readCsvFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	// lineCount := 0

	allRecords, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return allRecords
}

// csvFirstColumnToSlice Returns the first column of the csv as a slice
func csvFirstColumnToSlice(csvSlice [][]string) []string {
	var allUrls []string
	for _, v := range csvSlice {
		allUrls = append(allUrls, v[0])
	}
	return allUrls
}
