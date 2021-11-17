package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите произвольное количество целых чисел. Для завершения ввода нажмите ENTER без ввода числа:")

	for {
		if scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				fmt.Println("Это не число! Повторите снова или нажмите ENTER чтобы завершить ввод")
				continue
			}
			inputNums = append(inputNums, num)
		}
	}

	fmt.Println("Введена последовательность чисел: ", inputNums)

	// Сортировка простыми обменами, сортировка пузырьком
	for range inputNums {
		for i := 0; i < len(inputNums)-1; i++ {
			if inputNums[i] > inputNums[i+1] {
				inputNums[i], inputNums[i+1] = inputNums[i+1], inputNums[i]
			}
		}
	}

	fmt.Println("Последовательность чисел отсортирована: ", inputNums)

}
