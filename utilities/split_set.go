// Functions

// - Split into training and test sets
// - Split into training, verification and test sets
// - Sample if set is too large
// - Bootstrap set

package utilities

import (
	"../."
	"log"
	"math/rand"
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
	Train *Dataset
	Test  *Dataset
	Split []float64
}

// Import data from CSV
func (d *Dataset) LoadCSV(filename string) bool {
	if data, err := godatamining.FromCSV(filename); err == nil {
		tmp := (*data)[1:]
		d.Data = &tmp
		d.Size = uint16(len(*data) - 1)
		return true
	}
	return false
}

// Shuffle data for training/test set building
func (d *Dataset) Shuffle() (output Dataset) {
	data := *(*d).Data
	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	output.Data = &data
	output.Size = (*d).Size
	return
}

// Split set into training/test set n times
func (d *Dataset) SplitSet(ratio string, n int) (output []TestSet, ok bool) {
	rsplit := strings.Split(ratio, ",")
	split := make([]float64, len(rsplit))
	for i, a := range rsplit {
		b, _ := strconv.Atoi(a)
		split[i] = float64(b) / float64(100)
	}

	output = make([]TestSet, n)
	for i := 0; i < n; i++ {
		temp := d.Shuffle()
		output[i].Split = split
		output[i].Train = new(Dataset)
		output[i].Test = new(Dataset)

		train := (*temp.Data)[:int(float64(temp.Size)*split[0])]
		test := (*temp.Data)[int(float64(temp.Size)*split[0])+1:]

		output[i].Train.Data = &train
		output[i].Train.Size = uint16(float64(temp.Size) * split[0])

		output[i].Test.Data = &test
		output[i].Test.Size = uint16(float64(temp.Size) * split[1])
		log.Println(temp.Size, output[i].Test.Size, output[i].Train.Size, test)
	}
	ok = true
	return
}

// Output built sets into flat files
func (d *Dataset) SaveSet(location string) bool {
	return true
}

func (d *Dataset) BootStrapSet(n int) bool {
	return true
}
