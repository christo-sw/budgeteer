package internal

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadAndCategoriseCSV(file string) {

}

func readCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Unable to read input file '%s': %s", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse file as CSV for '%s': %s", filePath, err)
	}

	return records
}
