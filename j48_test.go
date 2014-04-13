package godatamining

import (
	"log"
	"testing"
)

func initJ48() (*J48, bool) {
	csv_file := "datasets/lenses.csv"
	oner := new(J48)
	data, err := FromCSV(csv_file)
	if err != nil {
		return oner, false
	}
	oner.Data = data
	return oner, true
}

func TestGetFrequency(t *testing.T) {
	if j, ok := initJ48(); ok {
		if _, ok := j.GetFrequency("Tears", "normal"); ok {
			return
		}
	}
	t.Fail()
}

func TestGetInfo(t *testing.T) {
	// get entropy
	if j, ok := initJ48(); ok {
		if out, ok := j.GetInfo(j.Data, "Tears"); ok {
			log.Println(out)
			return
		}
	}
	t.Fail()

}

func TestGetInfox(t *testing.T) {
	if j, ok := initJ48(); ok {
		if out, ok := j.GetInfox("Tears"); ok {
			log.Println(out)
			return
		}
	}
	t.Fail()

}

func TestGetInfoGain(t *testing.T) {
	if j, ok := initJ48(); ok {
		if out, ok := j.GetInfoGain("Tears"); ok {
			log.Println(out)
			return
		}
	}
	t.Fail()
}

func TestGetInfoRatio(t *testing.T) {}
