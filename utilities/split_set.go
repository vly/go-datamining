// Functions

// - Split into training and test sets
// - Split into training, verification and test sets
// - Bootstrap set

package utilities

import (
	"../."
	"log"
	"rand"
	"strconv"
	"strings"
)

// represents a data set
type Dataset struct {
	Data *[][]string
	Size uint16
}

// training/test set struct
type TestSet struct {
	Test  *Dataset
	Train *Dataset
	Split []int
}

// Import data from CSV
func (d *Dataset) LoadCSV(filename string) bool {
	if data, err := godatamining.FromCSV(filename); err == nil {
		d.Data = data
		d.Size = uint16(len(*data))
		return true
	}
	return false
}

// Shuffle data for training/test set building
func (d *Dataset) Shuffle() (output Dataset) {
	data = *(*d).Data
	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	output.Data = data
	output.Size = (*d).Size
}

// Split set into training/test set n times
func (d *Dataset) SplitSet(ratio string, n int) (output *[]Dataset, ok bool) {
	rsplit := strings.Split(ratio, ",")
	total := 0.0
	output = make([]Dataset, n)
	for i = 0; i < n; i++ {
		output
	}
	for _, a := range rsplit {
		b, _ := strconv.Atoi(a)
		total += b / float64(100)
	}
	log.Println(total)
	return
}

// Output built sets into flat files
func (d *Dataset) SaveSet(location string) bool {
	return true
}

func (d *Dataset) BootStrapSet(n int) bool {
	return true
}
