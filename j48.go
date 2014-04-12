package godatamining

import (
	"math"
)

type J48 struct {
	Data *[][]string
}

func (j *J48) getKeyLoc(key string) (int, bool) {
	keys := (*j.Data)[0]
	for i := 0; i < len(keys); i++ {
		if keys[i] == key {
			return i, true
		}
	}
	return -1, false
}

func (j *J48) GetFrequency(class string, value string) (int, bool) {
	total := 0
	if loc, ok := j.getKeyLoc(class); ok {
		for i := 1; i < len(*j.Data); i++ {
			if (*j.Data)[i][loc] == value {
				total += 1
			}
		}
		return total, true
	}
	return total, false
}

func (j *J48) genSet(class string) (output []string) {
	if loc, ok := j.getKeyLoc(class); ok {
		temp := make(map[string]bool)
		for i := 1; i < len(*j.Data); i++ {
			temp[(*j.Data)[i][loc]] = true
		}

		for a, _ := range temp {
			output = append(output, a)
		}
	}
	return
}
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
	for i := 1; i < len(*j.Data); i++ {
		temp, _ := indexOf(set, (*j.Data)[i][loc])
		output[temp] = append(output[temp], (*j.Data)[i])
	}
	return &output
}

func (j *J48) GetInfo(class string) (float32, bool) {
	var total float32
	set := j.genSet(class)
	for _, a := range set {
		b, _ := j.GetFrequency(class, a)
		x := float32(b) / float32(len(*j.Data)-1)
		total += x * float32(math.Pow(float64(x), 2))
	}
	return total, true
}

func (j *J48) GetInfox(class string) (float32, bool) {
	var total float32
	tables := j.splitTable(class)
	for i := 0; i < len(*tables); i++ {
		if info, ok := j.GetInfo(class); ok {
			total += float32(len((*tables)[i])) / float32(len(*j.Data)) * info
		}
	}
	return total, true
}

func (j *J48) GetInfoGain(class string) (float32, bool) {
	a, _ := j.GetInfo(class)
	b, _ := j.GetInfox(class)
	return a - b, true
}
