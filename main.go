package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс случайных положительных чисел заданного размера
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) + 1 // Положительные числа от 1 до 1 000 000
	}
	return data
}

// maximum находит максимальное число в слайсе
func maximum(data []int) int {
	if len(data) == 0 {
		return 0 // Для пустого слайса возвращаем 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks находит максимальное значение, разделяя слайс на части и обрабатывая их параллельно
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		startIndex := chunkIndex * chunkSize
		endIndex := startIndex + chunkSize

		// Для последнего чанка берём все оставшиеся элементы
		if chunkIndex == CHUNKS-1 {
			endIndex = len(data)
		}

		go func(chunkIdx int, start, end int) {
			defer wg.Done()
			chunkMax := data[start]
			for i := start + 1; i < end; i++ {
				if data[i] > chunkMax {
					chunkMax = data[i]
				}
			}
			maxValues[chunkIdx] = chunkMax
		}(chunkIndex, startIndex, endIndex)
	}

	wg.Wait()

	// Находим максимум среди максимальных значений чанков
	finalMax := maxValues[0]
	for i := 1; i < CHUNKS; i++ {
		if maxValues[i] > finalMax {
			finalMax = maxValues[i]
		}
	}
	return finalMax
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	startTime := time.Now()
	data := generateRandomElements(SIZE)
	generationTime := time.Since(startTime).Microseconds()
	fmt.Printf("Время генерации: %d мкс\n", generationTime)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime = time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxParallel, elapsedParallel)
}
