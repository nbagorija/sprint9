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

// generateRandomElements генерирует слайс случайных положительных целых чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) // Генерирует числа до 1 миллиона
	}
	return data
}

// maximum возвращает максимальное число в слайсе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks возвращает максимальное число, используя несколько горутин.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS // Деление с округлением вверх
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		if start < len(data) {
			wg.Add(1)
			go func(chunk []int, index int) {
				defer wg.Done()
				if len(chunk) > 0 {
					maxValues[index] = maximum(chunk)
				}
			}(data[start:end], i)
		}
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxParallel, elapsedParallel)
}
