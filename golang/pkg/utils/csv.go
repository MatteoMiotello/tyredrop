package utils

import (
	"encoding/csv"
	"os"
)

func CsvReadFile(filePath string, delimiters ...rune) ([][]string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	if len(delimiters) > 0 {
		reader.Comma = delimiters[0]
	}

	return reader.ReadAll()
}
