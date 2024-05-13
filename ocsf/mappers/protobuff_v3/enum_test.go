package protobuff_v3

import (
	"sort"
	"testing"
)

func TestValueSorter(t *testing.T) {
	values := []*EnumValue{
		{Name: "UNKNOWN", Value: 0},
		{Name: "B", Value: 2},
		{Name: "A", Value: 1},
		{Name: "C", Value: 3},
		{Name: "D", Value: 4},
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].Value < values[j].Value
	})

	expected := []string{
		"UNKNOWN",
		"A",
		"B",
		"C",
		"D",
	}

	for i, v := range values {
		if v.Name != expected[i] {
			t.Errorf("Expected %s, but got %s", expected[i], v.Name)
		}
	}
}
