package gorecommend

import (
	"fmt"
	"math"
)

// GetDistance is a basic example of a Euclidean distance measurement function
// based on a mock similarity matrix
func GetDistance(data map[string]map[string]int, key1 string, key2 string) (result float64, ok bool) {
	if value1, ok := data[key1]; ok {
		if value2, ok := data[key2]; ok {
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
					total += math.Pow(float64(data[key1][datum]-data[key2][datum]), 2)
				}
				return 1 / (1 + math.Sqrt(total)), true
			}
		}
	}
	return 0, false
}

type Distance struct {
	From     string
	To       string
	Distance float64
}

func (d *Distance) ToString() (output string) {
	return fmt.Sprintf("%s -> %s: %f", d.From, d.To, d.Distance)
}

// GetDistances generates all distances
func GetDistances(data map[string]map[string]int) ([]*Distance, bool) {
	var output []*Distance
	for k, _ := range data {
		for k2, _ := range data {
			if k != k2 {
				if temp, ok := GetDistance(data, k, k2); ok {
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
