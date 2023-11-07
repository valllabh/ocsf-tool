package protobuff_v3

import (
	"sort"
	"testing"
)

func TestValueSorter(t *testing.T) {
	values := []*EnumValue{
		{Name: "UNKNOWN"},
		{Name: "B"},
		{Name: "A"},
		{Name: "C"},
		{Name: "D"},
		{Name: "E"},
		{Name: "F"},
		{Name: "G"},
		{Name: "H"},
		{Name: "I"},
		{Name: "J"},
		{Name: "K"},
		{Name: "L"},
		{Name: "M"},
		{Name: "N"},
		{Name: "O"},
		{Name: "P"},
		{Name: "Q"},
		{Name: "R"},
		{Name: "S"},
		{Name: "T"},
		{Name: "U"},
		{Name: "V"},
		{Name: "W"},
		{Name: "X"},
		{Name: "Y"},
		{Name: "Z"},
	}

	sort.Slice(values, func(i, j int) bool {
		return valueSorter(values, i, j)
	})

	expected := []string{
		"UNKNOWN",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}

	for i, v := range values {
		if v.Name != expected[i] {
			t.Errorf("Expected %s, but got %s", expected[i], v.Name)
		}
	}
}
