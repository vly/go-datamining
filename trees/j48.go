package godatamining

import (
	"log"
	"math"
)

type J48 struct {
	Data *[][]string
	Keys *[]string
}

// Represents a given attribute
type Attribute struct {
	Rows  map[string]*Value
	Total int
}

// Represents frequency against a given class attribute
type Value struct {
	Data  map[string]int
	Total int
}

// Find the index of a given column header
func (j *J48) getKeyLoc(key string) (int, bool) {
	keys := (*j.Keys)
	for i := 0; i < len(keys); i++ {
		if keys[i] == key {
			return i, true
		}
	}
	return -1, false
}

// Calculate the entropy of a given row
func (j *J48) calcEntropy(row *Value) (float64, bool) {
	var total float64 = 0.0
	for a, b := range row.Data {
		prob := float64(b) / float64(row.Total)
		total -= prob * math.Log2(prob)
		log.Println(a, ": ", b, prob, total, row.Total)
	}
	return total, true
}

// Get Entropy using the frequency tables of class and predictor attributes
func (j *J48) getEntropy(attr *Attribute) (float64, bool) {
	var total float64 = 0.0
	for a, b := range attr.Rows {
		prob := float64(b.Total) / float64(attr.Total)
		ent, _ := j.calcEntropy(b)
		log.Println("getEntropy", a, prob, ent, b.Total, attr.Total)
		total += prob * ent
	}
	log.Println("total: ", total)
	return total, true
}

// Get frequency table of a single attribute (class)
func (j *J48) GetFrequency(data *[][]string, class string) (attr *Attribute, okay bool) {
	if cLoc, ok := j.getKeyLoc(class); ok {
		attr = new(Attribute)
		attr.Rows = make(map[string]*Value)
		attr.Total = len(*data)
		log.Println(attr.Total)
		okay = true
		for i := 0; i < attr.Total; i++ {
			if _, ok := attr.Rows[class]; ok {
				attr.Rows[class].Total += 1
				if _, ok := attr.Rows[class].Data[(*data)[i][cLoc]]; !ok {
					attr.Rows[class].Data[(*data)[i][cLoc]] = 1
				} else {
					attr.Rows[class].Data[(*data)[i][cLoc]] += 1
				}
			} else {
				t := new(Value)
				t.Total = 1
				t.Data = make(map[string]int)
				t.Data[(*data)[i][cLoc]] = 1
				attr.Rows[class] = t
			}
		}
	} else {
		okay = false
	}
	return
}

// As above but produces as DataFrame-esk slice
func (j *J48) getFrequencyX2(class string, predictor string) (attr *Attribute, rOk bool) {
	if tLoc, ok := j.getKeyLoc(predictor); ok {
		if cLoc, ok := j.getKeyLoc(class); ok {
			// classSet := j.genSet(class)
			attr = new(Attribute)
			rOk = true
			attr.Rows = make(map[string]*Value)
			attr.Total = len(*j.Data) - 1
			for i := 0; i < len(*j.Data); i++ {
				if _, ok := attr.Rows[(*j.Data)[i][tLoc]]; ok {
					attr.Rows[(*j.Data)[i][tLoc]].Total += 1
					if _, ok := attr.Rows[(*j.Data)[i][tLoc]].Data[(*j.Data)[i][cLoc]]; !ok {
						attr.Rows[(*j.Data)[i][tLoc]].Data[(*j.Data)[i][cLoc]] = 1
					} else {
						attr.Rows[(*j.Data)[i][tLoc]].Data[(*j.Data)[i][cLoc]] += 1
					}
				} else {
					t := new(Value)
					t.Total = 1
					t.Data = make(map[string]int)
					t.Data[(*j.Data)[i][cLoc]] = 1
					attr.Rows[(*j.Data)[i][tLoc]] = t
				}
			}
		} else {
			rOk = false
		}
	}
	return
}

func (j *J48) getInfo(data *[][]string, predictor string, class string) float64 {
	total := 0.0
	out := j.splitTable(predictor)
	for a, b := range *out {
		if attr, ok := j.GetFrequency(&b, class); ok {
			log.Println("dataset", a, b)
			if tmp, ok := j.getEntropy(attr); ok {
				total += float64(attr.Total) / float64(len(*j.Data)) * tmp
			}
		}
	}
	log.Println("getInfo", total)
	return total
}

// func (j *J48) GetFrequency(class string, value string) (int, bool) {
// 	total := 0
// 	if loc, ok := j.getKeyLoc(class); ok {
// 		for i := 1; i < len(*j.Data); i++ {
// 			if (*j.Data)[i][loc] == value {
// 				total += 1
// 			}
// 		}
// 		return total, true
// 	}
// 	return total, false
// }

// Create a set out of a given column
func (j *J48) genSet(class string) (output []string) {
	if loc, ok := j.getKeyLoc(class); ok {
		temp := make(map[string]bool)
		for i := 0; i < len(*j.Data); i++ {
			temp[(*j.Data)[i][loc]] = true
		}

		for a, _ := range temp {
			output = append(output, a)
		}
	}
	return
}

// Grab location of an item in a set
func indexOf(set []string, key string) (int, bool) {
	for i, a := range set {
		if a == key {
			return i, true
		}
	}
	return -1, false
}

// splitTable splits data on a given string
func (j *J48) splitTable(class string) *[][][]string {
	set := j.genSet(class)
	loc, _ := j.getKeyLoc(class)
	output := make([][][]string, len(set))
	for i := 0; i < len(*j.Data); i++ {
		temp, _ := indexOf(set, (*j.Data)[i][loc])
		output[temp] = append(output[temp], (*j.Data)[i])
	}
	return &output
}

// func (j *J48) GetInfo(table *[][]string, class string) (float64, bool) {
// 	var total float64
// 	set := j.genSet(class)
// 	for _, a := range set {
// 		b, _ := j.GetFrequency(class, a)
// 		x := float64(b) / float64(len(*table)-1)
// 		total += x * float64(math.Pow(float64(x), 2))
// 	}
// 	return -total, true
// }

// func (j *J48) GetInfox(class string) (float64, bool) {
// 	var total float64
// 	tables := j.splitTable(class)
// 	for i := 0; i < len(*tables); i++ {
// 		if info, ok := j.GetInfo(&(*tables)[i], class); ok {
// 			total += float64(len((*tables)[i])) / float64(len(*j.Data)) * info
// 		}
// 	}
// 	return total, true
// }

func (j *J48) getInfoGain(data *[][]string, predictor string, class string) (float64, bool) {
	output := 0.0
	info := j.getInfo(data, predictor, class)
	if attr, ok := j.GetFrequency(j.Data, class); ok {
		if orig, ok := j.getEntropy(attr); ok {
			log.Println(orig, info, orig-info)
			return output, true
		}
	}
	return output, false
}
