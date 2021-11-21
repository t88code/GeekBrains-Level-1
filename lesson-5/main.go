package main

import (
	"fmt"
	"time"
)

func fibonacciWithCache(n uint64, cache map[uint64]uint64) uint64 {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if cache == nil {
		return fibonacciWithCache(n-1, nil) + fibonacciWithCache(n-2, nil)
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = fibonacciWithCache(n-1, cache) + fibonacciWithCache(n-2, cache)
	return cache[n]

}

func fibonacciWithoutCache(n uint64) uint64 {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciWithoutCache(n-1) + fibonacciWithoutCache(n-2)

}

func main() {

	cache := make(map[uint64]uint64)
	var n uint64

	for {
		fmt.Print("Введите номер элемента ряда Фибоначчи: ")
		if _, err := fmt.Scanln(&n); err != nil {
			fmt.Println("Введено неверное значение")
			continue
		}
		startTimer := time.Now()
		fmt.Println("Запущена функция с кешем, кеш == nil")
		fmt.Println("Результат:", fibonacciWithCache(n, nil))
		fmt.Println("Время выполнения:", time.Now().Sub(startTimer))

		startTimer = time.Now()
		fmt.Println("Запущена функция с кешем, кеш != nil:")
		fmt.Println("Результат:", fibonacciWithCache(n, cache))
		fmt.Println("Время выполнения:", time.Now().Sub(startTimer))

		startTimer = time.Now()
		fmt.Println("Запущена функция без кеша:")
		fmt.Println("Результат:", fibonacciWithoutCache(n))
		fmt.Println("Время выполнения:", time.Now().Sub(startTimer))
	}

}
