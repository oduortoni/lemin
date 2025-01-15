package utils

import "testing"

func TestSliceContainsString(t *testing.T) {
	table := []struct {
		Slice  []string
		Str    string
		Expect bool
	}{
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "charlote",
			Expect: true,
		},
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "non-existent",
			Expect: false,
		},
		{
			Slice:  []string{},
			Str:    "anything",
			Expect: false,
		},
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "",
			Expect: false,
		},
	}

	for _, entry := range table {
		got := SliceContainsString(entry.Slice, entry.Str)
		if got != entry.Expect {
			t.Errorf("expected %v, but got %v\n", entry.Expect, got)
		}
	}
}
