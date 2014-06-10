package utilities

import (
	"testing"
)

func TestSplitSet(t *testing.T) {
	// FromCSV("asdf")
	output := new(Dataset)
	if ok := output.LoadCSV("../datasets/lenses.csv"); ok {
		output.SplitSet("33,66", 1)
		return
	}
	t.Fail()
}
