package utils

import (
	"testing"

	"lem-in/vars"
)

func TestProcessNumberOfAnts(t *testing.T) {
	restore := vars.AntsNumber

	table := []struct {
		Line    string
		Expect  func(error) bool
		Message string
	}{
		{
			Line: "90",
			Expect: func(err error) bool {
				return err == nil
			},
			Message: "Expected a nil error",
		},
		{
			Line: "Hello",
			Expect: func(err error) bool {
				return err != nil
			},
			Message: "Expected a non-nil error",
		},
	}

	for _, entry := range table {
		err := ProcessNumberOfAnts(entry.Line)
		if !entry.Expect(err) {
			t.Errorf("%s\n", entry.Message)
			vars.AntsNumber = restore
		}
	}
	vars.AntsNumber = restore
}
