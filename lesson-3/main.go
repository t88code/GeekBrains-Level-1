package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	for {
		if err := calculator(); err != nil {
			fmt.Println(err.Error())
			fmt.Println("Попробуйте снова!")
		}
	}

}

func calculator() error {

	var res float64
	var op string

	fmt.Print("Введите операцию на калькуляторе (+, -, *, /, sin, cos, tan, ln - натуральный логарифм, log - десятичный логарифм, sqrt - корень квадратный, pow - возведение в степень, ! - факториал, exit - выход): ")
	if _, err := fmt.Scanln(&op); err != nil {
		return fmt.Errorf("введено неверное значение")
	}

	if op == "+" || op == "-" || op == "*" || op == "/" {

		var a, b float64

		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scanln(&a); err != nil {
			return fmt.Errorf("введено неверное значение")
		}

		fmt.Print("Введите второе число: ")
		if _, err := fmt.Scanln(&b); err != nil {
			return fmt.Errorf("введено неверное значение")
		}

		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				return fmt.Errorf("делитель не может быть равен нулю")
			}
			res = a / b
		}

	} else if op == "sin" || op == "cos" || op == "tan" {

		var rad float64

		fmt.Print("Введите угол в радианах: ")
		if _, err := fmt.Scanln(&rad); err != nil {
			return fmt.Errorf("введено неверное значение")
		}

		switch op {
		case "sin":
			res = math.Sin(rad)
		case "cos":
			res = math.Cos(rad)
		case "tan":
			res = math.Tan(rad)
		}

	} else if op == "ln" || op == "log" {

		var a float64

		fmt.Print("Введите любое число больше нуля: ")
		if _, err := fmt.Scanln(&a); err != nil {
			return fmt.Errorf("введено неверное значение")
		}
		if a <= 0 {
			return fmt.Errorf("значение должно быть больше нуля")
		}

		switch op {
		case "ln":
			res = math.Log(a)
		case "log":
			res = math.Log10(a)
		}

	} else if op == "sqrt" {

		var a float64

		fmt.Print("Введите любое число большее или равное нулю: ")
		if _, err := fmt.Scanln(&a); err != nil {
			return fmt.Errorf("введено неверное значение")
		}
		if a < 0 {
			return fmt.Errorf("значение должно быть больше и равно нулю")
		}

		res = math.Sqrt(a)

	} else if op == "pow" {

		var x, y float64

		fmt.Print("Введите число: ")
		if _, err := fmt.Scanln(&x); err != nil {
			return fmt.Errorf("введено неверное значение")
		}

		fmt.Print("Введите степень: ")
		if _, err := fmt.Scanln(&y); err != nil {
			return fmt.Errorf("введено неверное значение")
		}

		res = math.Pow(x, y)

	} else if op == "!" {

		var a, s uint

		fmt.Print("Введите любое натуральное число больше нуля: ")
		if _, err := fmt.Scanln(&a); err != nil {
			return fmt.Errorf("введено неверное значение")
		}
		if a <= 0 {
			return fmt.Errorf("значение должно быть больше нуля")
		}

		for s = 1; a > 0; a-- {
			s = s * a
		}
		res = float64(s)

	} else if op == "exit" {

		os.Exit(1)

	} else {

		return fmt.Errorf("операция выбрана неверно")

	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
	return nil
}
