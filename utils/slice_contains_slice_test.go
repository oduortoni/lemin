package utils

import "testing"

func TestSliceContainsSlice(t *testing.T) {
	table := []struct {
		Haystack []string
		Kneedle  []string
		Expect   bool
	}{
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{"charlote"},
			Expect:   true,
		},
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{"non-existent"},
			Expect:   false,
		},
		{
			Haystack: []string{},
			Kneedle:  []string{"anything"},
			Expect:   false,
		},
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{""},
			Expect:   false,
		},
	}

	for _, entry := range table {
		got := SliceContainsSlice(entry.Haystack, entry.Kneedle)
		if got != entry.Expect {
			t.Errorf("expected %v, but got %v\n", entry.Expect, got)
		}
	}
}
