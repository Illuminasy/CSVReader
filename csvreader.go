// Package csvreader ...
// Maintainer Illuminasy
// Packege to extract csv contents.
package csvreader

import (
	"encoding/csv"
	"os"
)

// CSVRow column values for a  single line in csv
type CSVRow []string

// CSVRows lines from csv file
type CSVRows []CSVRow

// ExtractContents extracts array of lines from csv file
func ExtractContents(fileName string) ([][]string, error) {
	return readCSVFile(fileName)
}

// readCSVFile Reads csv file and returns array of lines all column values.
// Returns error if file is not found or fails to read the file.
func readCSVFile(fileName string) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	// Read File into a Variable
	return csv.NewReader(f).ReadAll()
}
