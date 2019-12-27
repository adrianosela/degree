package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	parseFile("degree.csv", func(record []string) {
		fmt.Println(record[2])
	})
}

// parseFile operates on a CSV file by ignoring the first
// row (header) and running the handler for every other row
func parseFile(path string, handler func([]string)) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open csv file: %s", err)
	}
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return fmt.Errorf("could not read csv file header: %s", err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read csv row: %s", err)
		}
		handler(record)
	}
	return nil
}
