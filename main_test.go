package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"положительный размер", 10, 10},
		{"нулевой размер", 0, 0},
		{"отрицательный размер", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			assert.Equal(t, tt.expected, len(result),
				"для размера %d ожидался слайс длиной %d, но получен %d",
				tt.size, tt.expected, len(result))
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"несколько элементов", []int{1, 5, 3, 9, 2}, 9},
		{"одинаковые элементы", []int{7, 7, 7}, 7},
		{"убывающая последовательность", []int{9, 7, 5, 3}, 9},
		{"возрастающая последовательность", []int{1, 3, 5, 7, 9}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result,
				"для слайса %v ожидался максимум %d, но получено %d",
				tt.input, tt.expected, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"маленький слайс (меньше чанков)", []int{1, 3, 2}, 3},
		{"нормальный слайс", []int{1, 10, 3, 8, 5, 12, 7, 6}, 12},
		{"все элементы одинаковые", []int{5, 5, 5, 5}, 5},
		{"максимум в начале", []int{15, 1, 2, 3}, 15},
		{"максимум в конце", []int{1, 2, 3, 20}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			assert.Equal(t, tt.expected, result,
				"для слайса %v ожидался максимум %d, но получено %d",
				tt.input, tt.expected, result)
		})
	}
}
