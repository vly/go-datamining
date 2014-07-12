package godatamining

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Result struct {
	Key        string
	TotalError float32
	Rules      *map[string]float32
}

func FromCSV(filename string) (records *[][]string, err error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	lineCount := 0
	records = new([][]string)
	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			break
		}
		*records = append(*records, record)
		lineCount += 1
	}
	return
}

func ToCSV(filename string, header *[]string) bool {
	csvFile, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	return ok
}
