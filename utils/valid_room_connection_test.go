package utils

import "testing"

func TestValidRoomConnection(t *testing.T) {
	table := []struct {
		Input  string
		Expect bool
	}{
		{
			Input:  "4-7",
			Expect: true,
		},
		{
			Input:  "4 7",
			Expect: false,
		},
		{
			Input:  "start-7",
			Expect: true,
		},
		{
			Input:  "4-end",
			Expect: true,
		},
		{
			Input:  "4-",
			Expect: false,
		},
		{
			Input:  "jakarta-kisumu",
			Expect: true,
		},
	}

	for _, entry := range table {
		got := ValidRoomConnection(entry.Input)
		if got != entry.Expect {
			t.Errorf("All valid room connections must have a '-' between them, but got %s\n", entry.Input)
		}
	}
}
