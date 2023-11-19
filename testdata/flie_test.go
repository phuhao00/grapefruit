package testdata

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
	if result == 5 {
		fmt.Println(result)
	}
}
