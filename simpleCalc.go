package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите арифметическую операцию: ")
	fmt.Scanln(&op)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b != 0 {
			res = a / b
		} else {
			fmt.Println("Делить на ноль нельзя")
			os.Exit(2)
		}
	case "^":
		res = float64(math.Pow(a, b))
	case "!":
		if a >= 0 {
			var y int = int(a)
			res = float64(factorial(y))
		} else {
			fmt.Println("Введите положительное число")
			os.Exit(3)
		}

	default:
		fmt.Println("Операция вебрана не верно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)

}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
