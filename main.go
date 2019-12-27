package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reqs := make(map[string]int)

	operateCSV("requirements.csv", func(record []string) error {
		credits, err := strconv.Atoi(record[1])
		if err != nil {
			return fmt.Errorf("could not convert credits to integer: %s", err)
		}
		reqs[record[0]] = credits
		return nil
	})

	operateCSV("degree.csv", func(record []string) error {
		course := record[2]
		requirement := record[4]
		credits, err := strconv.Atoi(record[5])
		if err != nil {
			return fmt.Errorf("could not convert credits to integer: %s", err)
		}

		if _, ok := reqs[requirement]; ok {
			reqs[requirement] -= credits
			fmt.Printf("course %s applied to requirement: %s\n", course, requirement)
			if reqs[requirement] <= 0 {
				fmt.Printf("satisfied requirement: %s!\n", requirement)
				delete(reqs, requirement)
			}
		}
		return nil
	})

	if len(reqs) == 0 {
		fmt.Println("ALL REQUIREMENTS SATISFIED!")
	} else {
		for k, v := range reqs {
			fmt.Printf("UNSATISFIED REQUIREMENT: %s, MISSING %d credits\n", k, v)
		}
	}
}

// operateCSV operates on a CSV file by ignoring the first
// row (header) and running the handler for every other row
func operateCSV(path string, handler func([]string) error) error {
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
		if err = handler(record); err != nil {
			return fmt.Errorf("record handler errored: %s", err)
		}
	}
	return nil
}
