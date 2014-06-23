package godatamining

import (
	"log"
	"testing"
)

func initZeroR() (*ZeroR, bool) {
	csv_file := "datasets/lenses.csv"
	z := new(ZeroR)
	data, err := FromCSV(csv_file)
	if err != nil {
		return z, false
	}
	z.Data = data
	return z, true
}

func TestGetResult(t *testing.T) {
	z, ok := initZeroR()
	if !ok {
		t.Fail()
	}

	if rule, ok := z.GetResult("Lenses"); ok {
		log.Printf("ZeroR results: Class=%s, Rules=%v, Total error=%f, Accuracy=%f\n",
			rule.Key, *rule.Rules, rule.TotalError, (1 - rule.TotalError))
	} else {
		t.Fail()
	}

}
