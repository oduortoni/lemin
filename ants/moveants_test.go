package ants

import (
	"regexp"
	"strings"
	"testing"

	"lem-in/models"
)

type Printer struct {
	Buffer string
}

func (p *Printer) Println(args ...any) (n int, err error) {
	for _, arg := range args {
		if s, ok := arg.(string); ok {
			p.Buffer += s
		}
	}
	return len(p.Buffer), nil
}

func (p Printer) String() string {
	// fmt.Println(p.Buffer)
	return p.Buffer
}

// Test function to check output format with regex
func TestMoveAntsRegex(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"r1", "r2", "r10"},
			Ants:  []int{1, 2, 3, 4, 8, 12, 16, 20},
		},
		{
			Rooms: []string{"r1", "r3", "r5", "r7", "r8", "r10"},
			Ants:  []int{5, 9, 13, 17},
		},
		{
			Rooms: []string{"r1", "r4", "r6", "r10"},
			Ants:  []int{6, 10, 14, 18},
		},
		{
			Rooms: []string{"r1", "r9", "r10"},
			Ants:  []int{7, 11, 15, 19},
		},
	}

	// initialize the printer we will use to capture data that MoveAnts prints
	printer := Printer{
		Buffer: "",
	}

	// run the function
	MoveAnts(printer.Println, paths)

	// define the expected pattern and its regex
	// regex := `L\d+-\w+`
	regex := `L\d+-.+`
	re := regexp.MustCompile(regex)

	// split captured output into individual lines we can compare against the compiled regex
	lines := strings.Split(printer.String(), "\n")

	for _, line := range lines {
		if line != "" { // skip empty lines
			matched := re.Match([]byte(line))
			if !matched {
				t.Errorf("Line does not match the expected pattern: %s", line)
			}
		}
	}
}
