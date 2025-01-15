package utils

import (
	"os"
	"testing"
)

func TestHasStartAndEnd(t *testing.T) {
	file, err := os.CreateTemp("", "temp_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up after test

	_, err = file.WriteString("4\n##start\n0 0 3\n2 2 5\n3 4 0\n##end\n1 8 3\n0-2\n2-3\n3-1\n")
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// open/close to test if file can handle the open/read cycles
	err = file.Close() // Close the file after writing
	if err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	file, err = os.Open(file.Name()) // reopen the file for reading
	if err != nil {
		t.Fatalf("Failed to open file for testing: %v", err)
	}
	defer file.Close()

	isValid := HasStartAndEnd(file)
	if !isValid {
		t.Errorf("Expected the file to be valid, but got false")
	}
}
