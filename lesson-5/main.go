package main

import (
	"fmt"
	"time"
)

func fibonacci(n uint64, cache map[uint64]uint64) uint64 {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if cache[n] != 0 {
		return cache[n]
	}

	cache[n] = fibonacci(n-1, cache) + fibonacci(n-2, cache)
	return cache[n]

}

type fibonacciResult struct {
	// можно добавить историю запросов
	n     uint64
	cache map[uint64]uint64
}

func main() {

	f := fibonacciResult{
		cache: make(map[uint64]uint64),
	}

	for {
		fmt.Print("Введите номер элемента ряда Фибоначчи: ")
		if _, err := fmt.Scanln(&f.n); err != nil {
			fmt.Println("Введено неверное значение")
		}
		startTimer := time.Now()

		fmt.Println("Результат:", fibonacci(f.n, f.cache))
		fmt.Println("Время выполнения:", time.Now().Sub(startTimer))
	}

}
