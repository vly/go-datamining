package godatamining

import (
	"log"
	"testing"
)

func initOneR() (*OneR, bool) {
	csv_file := "datasets/lenses.csv"
	oner := new(OneR)
	if err := oner.FromCSV(csv_file); err != nil {
		return oner, false
	}
	return oner, true
}

func TestGetInstance(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetInstance("Prescription"); !ok {
		t.Fail()
	}

}

func TestGetRules(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetRules("Prescription"); !ok {
		t.Fail()
	}
}

func TestGetErrorRate(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetErrorRate("Prescription"); !ok {
		t.Fail()
	}
}

func TestGetBestRule(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if rule, ok := r.GetBestRule(); ok {
		log.Printf("Best rule: Class=%s, Rules=%v, Total error=%f\n",
			rule.Key, rule.Rules, rule.TotalError)
	} else {
		t.Fail()
	}
}
