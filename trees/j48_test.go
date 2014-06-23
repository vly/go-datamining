package godatamining

import (
	common "../."
	"log"
	"testing"
)

func initJ48() (*J48, bool) {
	csv_file := "../datasets/lenses_num.csv"
	oner := new(J48)
	data, err := common.FromCSV(csv_file)
	if err != nil {
		return oner, false
	}
	outData := (*data)[1:]
	oner.Data = &outData
	oner.Keys = &(*data)[0]
	log.Println(data)
	return oner, true
}

// Test Frequency iteration
func TestGetEntropy(t *testing.T) {
	if j, ok := initJ48(); ok {
		if attr, ok := j.getFrequencyX("Lenses", "Tears"); ok {
			j.getEntropy(attr)
			return
		}
	}
	t.Fail()
}

func TestGetFrequency(t *testing.T) {
	if j, ok := initJ48(); ok {
		if attr, ok := j.GetFrequency(j.Data, "Lenses"); ok {
			log.Println("TestGetFrequency")
			j.getEntropy(attr)
			return
		}
	}
	t.Fail()
}

// Test Frequency iteration
func TestGetFrequencyX(t *testing.T) {
	if j, ok := initJ48(); ok {
		if _, ok := j.getFrequencyX("Lenses", "Prescription"); ok {
			return
		}
	}
	t.Fail()
}

// Test split table on a given predictor attribute
func TestSplitTable(t *testing.T) {
	if j, ok := initJ48(); ok {
		if out := j.splitTable(j.Data, "Prescription"); len(*out) > 0 {
			return
		}
	}
	t.Fail()
}

func TestGetInfo(t *testing.T) {
	if j, ok := initJ48(); ok {
		j.getInfo(j.Data, "Age", "Lenses")
		return
	}
	t.Fail()
}

func TestGetInfoGain(t *testing.T) {
	if j, ok := initJ48(); ok {
		j.getInfoGain(j.Data, "Age", "Lenses")
		return
	}
	t.Fail()
}

func TestBuildTree(t *testing.T) {
	// get entropy
	if j, ok := initJ48(); ok {
		j.buildTree(j.Data, "Lenses")
		return
	}
	t.Fail()

}

// func TestGetInfox(t *testing.T) {
// 	if j, ok := initJ48(); ok {
// 		if out, ok := j.GetInfox("Tears"); ok {
// 			log.Println(out)
// 			return
// 		}
// 	}
// 	t.Fail()

// }

// func TestGetInfoGain(t *testing.T) {
// 	if j, ok := initJ48(); ok {
// 		if out, ok := j.GetInfoGain("Tears"); ok {
// 			log.Println(out)
// 			return
// 		}
// 	}
// 	t.Fail()
// }

// func TestGetInfoRatio(t *testing.T) {}
