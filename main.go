package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("Размер должен быть больше нуля")
		return nil
	}

	slice := make([]int, size)
	rand.Seed(10)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) + 1
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Println("Передан пустой слайс")
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	max := data[0] 
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		fmt.Println("Передан пустой слайс") 
		return 0
	}
	
	sliceOfMaxes := make([]int, CHUNKS)
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS // деление с округлением вверх

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := (i + 1) * chunkSize

		if end > len(data) {
			end = len(data) // чтобы не выходить за рамки слайса
		}

		if start >= end { // для случаев когда CHUNKS > len(data) для поздних пустых чанков
			wg.Done()
			continue
		}

		go func(s []int) {
			defer wg.Done()
			max := s[0]
			for _, v := range s[1:] {
				if v > max {
					max = v
				}
			}
			sliceOfMaxes[i] = max
		}(data[start:end])
	}
	wg.Wait()
	
	result := maximum(sliceOfMaxes)
	return result
}


func main() {
	fmt.Printf("Генерируем %d целых чисел\n\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice)
	duration := time.Since(start).Microseconds() 

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n\n", max, duration)
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS) 
	start = time.Now()
	max = maxChunks(slice)
	duration = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, duration)
}
