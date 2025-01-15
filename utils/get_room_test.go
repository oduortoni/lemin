package utils

import "testing"

func TestGetRoom(t *testing.T) {
	table := []struct {
		Line    string
		Expect  func(int, int, error) bool
		Message string
	}{
		{
			Line: "90 2 27",
			Expect: func(x, y int, err error) bool {
				return err == nil && x == 2 && y == 27
			},
			Message: "Expected a nil error",
		},
		{
			Line: "Hello",
			Expect: func(x, y int, err error) bool {
				return err != nil
			},
			Message: "Expected a non-nil error",
		},
	}

	for _, entry := range table {
		_, x, y, err := GetRoom(entry.Line)
		if !entry.Expect(x, y, err) {
			t.Errorf("%s\n", entry.Message)
		}
	}
}
