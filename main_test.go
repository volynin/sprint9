package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	// Тест с положительным размером
	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Ожидался слайс размером %d, но получен %d", size, len(result))
	}

	// Тест с нулевым размером
	empty := generateRandomElements(0)
	if len(empty) != 0 {
		t.Error("Ожидался пустой слайс для размера 0")
	}

	// Тест с отрицательным размером
	negative := generateRandomElements(-5)
	if len(negative) != 0 {
		t.Error("Ожидался пустой слайс для отрицательного размера")
	}
}

func TestMaximum(t *testing.T) {
	// Тест с пустым слайсом
	empty := []int{}
	if result := maximum(empty); result != 0 {
		t.Errorf("Для пустого слайса ожидался 0, но получено %d", result)
	}

	// Тест с одним элементом
	single := []int{42}
	if result := maximum(single); result != 42 {
		t.Errorf("Для слайса [42] ожидался 42, но получено %d", result)
	}

	// Тест с несколькими элементами
	multiple := []int{1, 5, 3, 9, 2}
	if result := maximum(multiple); result != 9 {
		t.Errorf("Для слайса [1,5,3,9,2] ожидался 9, но получено %d", result)
	}

	// Тест с одинаковыми элементами
	same := []int{7, 7, 7}
	if result := maximum(same); result != 7 {
		t.Errorf("Для слайса [7,7,7] ожидался 7, но получено %d", result)
	}
}

func TestMaxChunks(t *testing.T) {
	// Тест с пустым слайсом
	empty := []int{}
	if result := maxChunks(empty); result != 0 {
		t.Errorf("Для пустого слайса ожидался 0, но получено %d", result)
	}

	// Тест с одним элементом
	single := []int{42}
	if result := maxChunks(single); result != 42 {
		t.Errorf("Для слайса [42] ожидался 42, но получено %d", result)
	}

	// Тест с небольшим слайсом (меньше количества чанков)
	small := []int{1, 3, 2}
	if result := maxChunks(small); result != 3 {
		t.Errorf("Для слайса [1,3,2] ожидался 3, но получено %d", result)
	}

	// Тест с нормальным слайсом
	normal := []int{1, 10, 3, 8, 5, 12, 7, 6}
	if result := maxChunks(normal); result != 12 {
		t.Errorf("Для слайса [1,10,3,8,5,12,7,6] ожидался 12, но получено %d", result)
	}
}
