package godatamining

import (
	"log"
	"testing"
)

func initZeroR() (*ZeroR, bool) {
	csv_file := "datasets/lenses.csv"
	z := new(ZeroR)
	if err := z.FromCSV(csv_file); err != nil {
		return z, false
	}
	return z, true
}

func TestGetResult(t *testing.T) {
	z, ok := initZeroR()
	if !ok {
		t.Fail()
	}

	if rule, ok := z.GetResult("Lenses"); ok {
		log.Printf("ZeroR results: Class=%s, Rules=%v, Total error=%f\n",
			rule.Key, rule.Rules, rule.TotalError)
	} else {
		t.Fail()
	}

}
