// Functions

// - Split into training and test sets
// - Split into training, verification and test sets
// - Bootstrap set

package utilities

import (
	"../."
)

type Dataset struct {
	Data *[]string
	Size uint16
}

func (d *Dataset) SplitSet(ratio string, n int) bool {
	// FromCSV("asdf")
	godatamining.FromCSV("asdf")
	return true
}

func (d *Dataset) SaveSet(location string) bool {
	return true
}

func (d *Dataset) BootStrapSet(n int) bool {
	return true
}
