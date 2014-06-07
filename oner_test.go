package godatamining

import (
	"fmt"
	"log"
	"testing"
)

const class = "Lenses"

func initOneR() (*OneR, bool) {
	csv_file := "datasets/lenses.csv"
	oner := new(OneR)
	data, err := FromCSV(csv_file)
	if err != nil {
		return oner, false
	}
	oner.Data = data
	return oner, true
}

func TestGetInstance(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetInstance("Prescription", class); !ok {
		t.Fail()
	}

}

func TestGetRules(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetRules("Prescription", class); !ok {
		t.Fail()
	}
}

func TestGetErrorRate(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if _, ok := r.GetErrorRate("Tears", class); !ok {
		t.Fail()
	}
}

func TestGetBestRule(t *testing.T) {
	r, ok := initOneR()
	if !ok {
		t.Fail()
	}

	if rule, ok := r.GetBestRule(class); ok {
		tmp := ""
		for a, b := range *rule.Rules {
			tmp += fmt.Sprintf("If '%s' then ", a)
			for x, y := range b {
				tmp += fmt.Sprintf("'%s' (error: %f), ", x, y)
			}
		}
		log.Printf("Best rule: Class=%s, Attribute=%s, Rules=[%s] Total error=%f, Accuracy=%f\n",
			class, rule.Key, tmp, rule.TotalError, 1-rule.TotalError)
	} else {
		t.Fail()
	}
}
