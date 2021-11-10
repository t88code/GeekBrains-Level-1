package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var n int

	fmt.Println("Урок 2\n" +
		"Введите номер программы:\n" +
		"1 - Программа для вычисления площади прямоугольника\n" +
		"2 - Программа для вычисления диаметра и длины окружности по заданной площади круга\n" +
		"3 - Программа для определения сотен, десятков и единиц в трехзначном числе")

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка:", err.Error())
		os.Exit(1)
	}

	switch n {
	case 1:
		calcAreaRectangle()
	case 2:
		calcCircleParametrs()
	case 3:
		calcNumberRank()
	default:
		fmt.Println("Номер программы не определен")
	}
}

func calcAreaRectangle() {

	fmt.Println("Выбрана программа для вычисления площади прямоугольника")

	var a, b, s float64

	fmt.Print("Введите длину a: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка:", err.Error())
		os.Exit(1)
	}

	fmt.Print("Введите длину b: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка:", err.Error())
		os.Exit(1)
	}

	s = a * b

	fmt.Println("Площадь прямоугольника - ", s)
}

func calcCircleParametrs() {

	fmt.Println("Выбрана программа для вычисления диаметра и длины окружности по заданной площади круга")

	var d, l, s float64
	pi := math.Pi

	fmt.Print("Введите площадь окружности: ")
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Println("Ошибка:", err.Error())
		os.Exit(1)
	}

	d = 2 * math.Sqrt(s/pi)
	l = d * pi

	fmt.Println("Диаметр окружности - ", d)
	fmt.Println("Длина окружности - ", l)

}

func calcNumberRank() {

	fmt.Println("Выбрана программа для определения сотен, десятков и единиц в трехзначном числе")

	var x, s, d, e int

	fmt.Print("Введите трехзначное число: ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Ошибка:", err.Error())
		os.Exit(1)
	}

	if x < 100 || x > 999 {
		fmt.Println("Необходимо ввести трехзначное число")
		os.Exit(1)
	}

	s = x / 100
	d = (x - 100*s) / 10
	e = x - 100*s - 10*d

	fmt.Println("Сотен - ", s)
	fmt.Println("Десятков - ", d)
	fmt.Println("Единиц - ", e)

}
