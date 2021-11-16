package main

import (
	"fmt"
	"math"
)

func main() {

	var n int

	fmt.Println("Урок 2\n" +
		"Введите номер программы:\n" +
		"1 - Программа для вычисления площади прямоугольника\n" +
		"2 - Программа для вычисления диаметра и длины окружности по заданной площади круга\n" +
		"3 - Программа для определения сотен, десятков и единиц в трехзначном числе")

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Ошибка: %s", err.Error())
	}

	switch n {
	case 1:
		if err := calcAreaRectangle(); err != nil {
			fmt.Println(err.Error())
		}
	case 2:
		if err := calcCircleParametrs(); err != nil {
			fmt.Println(err.Error())
		}
	case 3:
		if err := calcNumberRank(); err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("Номер программы не определен")
	}
}

func calcAreaRectangle() error {

	fmt.Println("Выбрана программа для вычисления площади прямоугольника")

	var a, b, s float64

	fmt.Print("Введите длину a: ")
	if _, err := fmt.Scan(&a); err != nil {
		return fmt.Errorf("Ошибка: %s", err.Error())
	}

	fmt.Print("Введите длину b: ")
	if _, err := fmt.Scan(&b); err != nil {
		return fmt.Errorf("Ошибка: %s", err.Error())
	}

	s = a * b

	fmt.Println("Площадь прямоугольника - ", s)
	return nil
}

func calcCircleParametrs() error {

	fmt.Println("Выбрана программа для вычисления диаметра и длины окружности по заданной площади круга")

	var d, l, s float64

	fmt.Print("Введите площадь окружности: ")
	if _, err := fmt.Scan(&s); err != nil {
		return fmt.Errorf("Ошибка: %s", err.Error())
	}

	d = 2 * math.Sqrt(s/math.Pi)
	l = d * math.Pi

	fmt.Println("Диаметр окружности - ", d)
	fmt.Println("Длина окружности - ", l)
	return nil

}

func calcNumberRank() error {

	fmt.Println("Выбрана программа для определения сотен, десятков и единиц в трехзначном числе")

	var x, s, d, e int

	fmt.Print("Введите трехзначное число: ")
	if _, err := fmt.Scan(&x); err != nil {
		return fmt.Errorf("Ошибка: %s", err.Error())
	}

	if x < 100 || x > 999 {
		return fmt.Errorf("Необходимо ввести трехзначное число")
	}

	s = x / 100
	d = (x - 100*s) / 10
	e = x - 100*s - 10*d

	fmt.Println("Сотен - ", s)
	fmt.Println("Десятков - ", d)
	fmt.Println("Единиц - ", e)
	return nil
}
