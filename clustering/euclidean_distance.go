package godatamining

import (
	"fmt"
	"math"
)

type Euclidean struct {
	Data map[string]map[string]int
}

type Distance struct {
	From     string
	To       string
	Distance float64
}

func (d *Distance) ToString() (output string) {
	return fmt.Sprintf("%s -> %s: %f", d.From, d.To, d.Distance)
}

// GetDistance is a basic example of a Euclidean distance measurement function
// based on a mock similarity matrix
func (e *Euclidean) GetDistance(key1 string, key2 string) (result float64, ok bool) {
	if value1, ok := e.Data[key1]; ok {
		if value2, ok := e.Data[key2]; ok {
			var output []string
			for k, _ := range value1 {
				if _, ok := value2[k]; ok {
					output = append(output, k)
				}
			}
			// if there are similarities, lets check them
			if len(output) != 0 {
				total := float64(0)
				for _, datum := range output {
					total += math.Pow(float64(e.Data[key1][datum]-e.Data[key2][datum]), 2)
				}
				return 1 / (1 + math.Sqrt(total)), true
			}
		}
	}
	return 0, false
}

// GetDistances generates all distances
func (e *Euclidean) GetDistances() ([]*Distance, bool) {
	var output []*Distance
	for k, _ := range e.Data {
		for k2, _ := range e.Data {
			if k != k2 {
				if temp, ok := e.GetDistance(k, k2); ok {
					output = append(output, &Distance{k, k2, temp})
				}
			}
		}
	}

	if len(output) > 0 {
		return output, true
	}
	return output, false
}
