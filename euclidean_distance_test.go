package gorecommend

import (
	"log"
	"testing"
)

func TestGetDistance(t *testing.T) {
	log.Println("Running GetDistance test")
	data := map[string]map[string]int{
		"p1": map[string]int{
			"item1": 83,
			"item2": 44,
			"item4": 12,
		},
		"p2": map[string]int{
			"item1": 82,
			"item3": 90,
			"item4": 29,
		},
		"p3": map[string]int{
			"item1": 49,
			"item3": 33,
		},
		"p4": map[string]int{
			"item1": 82,
			"item3": 91,
			"item4": 28,
		},
	}
	if _, ok := GetDistance(data, "p1", "p2"); !ok {
		t.Fatalf("Error: %s \n", "Failed to get distance.")
	}
}

func TestGetDistances(t *testing.T) {
	data := map[string]map[string]int{
		"p1": map[string]int{
			"item1": 83,
			"item2": 44,
			"item4": 12,
		},
		"p2": map[string]int{
			"item1": 82,
			"item3": 90,
			"item4": 29,
		},
		"p3": map[string]int{
			"item1": 49,
			"item3": 33,
		},
		"p4": map[string]int{
			"item1": 82,
			"item3": 91,
			"item4": 28,
		},
	}
	output, ok := GetDistances(data)
	if !ok {
		t.Fatalf("Error: %s \n", "Failed to get distances.")
	}
	for a, b := range output {
		log.Printf("%d: %s", a, b.ToString())
	}

}
