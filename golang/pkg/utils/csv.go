package utils

import (
	"encoding/csv"
	"os"
)

func CsvReadFile(filePath string, opts ...interface{}) ([][]string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	if len(opts) > 0 {
		r := opts[0]

		reader.Comma = r.(rune)
	}

	return reader.ReadAll()
}
