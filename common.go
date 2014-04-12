package godatamining

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

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
