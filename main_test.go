package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(0)
	if len(slice) != 0 {
		t.Errorf("Expected length 0, got %d", len(slice))
	}

	size := 10
	slice = generateRandomElements(size)
	if len(slice) != size {
		t.Errorf("Expected length %d, got %d", size, len(slice))
	}

	for i, val := range slice {
		if val < 0 {
			t.Errorf("Element at index %d is negative: %d", i, val)
		}
	}
}

func TestMaximum(t *testing.T) {
	if max := maximum([]int{}); max != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", max)
	}

	if max := maximum([]int{42}); max != 42 {
		t.Errorf("Expected 42, got %d", max)
	}

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{10, 5, 20, 15}, 20},
		{[]int{100, 100, 100}, 100},
		{[]int{5, 1, 2, 3, 4}, 5},
	}

	for _, tc := range testCases {
		if max := maximum(tc.input); max != tc.expected {
			t.Errorf("For input %v expected %d, got %d", tc.input, tc.expected, max)
		}
	}
}

func TestMaxChunks(t *testing.T) {
	if max := maxChunks([]int{}); max != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", max)
	}

	smallSlice := []int{1, 2, 3, 4, 5}
	if max := maxChunks(smallSlice); max != 5 {
		t.Errorf("Expected 5, got %d", max)
	}

	largeSlice := make([]int, 100)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	if max := maxChunks(largeSlice); max != 99 {
		t.Errorf("Expected 99, got %d", max)
	}

	data := generateRandomElements(1000)
	max1 := maximum(data)
	max2 := maxChunks(data)
	if max1 != max2 {
		t.Errorf("Results don't match: maximum=%d, maxChunks=%d", max1, max2)
	}
}
