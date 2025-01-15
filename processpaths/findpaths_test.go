package processpaths

import (
	"fmt"
	"testing"

	"lem-in/vars"
)

func TestFindpaths(t *testing.T) {
	colony := vars.Colony
	allPaths := vars.AllPaths

	vars.Colony = map[string][]string{
		"0":     {"o", "start"},
		"A":     {"c", "h"},
		"E":     {"a", "t"},
		"a":     {"m", "E"},
		"c":     {"A", "k"},
		"e":     {"n", "end"},
		"end":   {"k", "m", "e"},
		"h":     {"start", "A", "n"},
		"k":     {"end", "c"},
		"m":     {"a", "end", "n"},
		"n":     {"e", "o", "m", "h"},
		"o":     {"0", "n"},
		"start": {"t", "h", "0"},
		"t":     {"start", "E"},
	}

	// fmt.Println("B4: ", vars.AllPaths)
	FindPaths("start", "end")

	if len(vars.AllPaths) <= 0 {
		// restore global variables
		vars.Colony = colony
		vars.AllPaths = allPaths
		t.Errorf("Expected the paths to be full")
	} else {
		// restore global variables
		fmt.Println(vars.AllPaths)
		vars.Colony = colony
		vars.AllPaths = allPaths
	}
}
