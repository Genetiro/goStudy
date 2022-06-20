package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)

}
func main() {
	var n int
	fmt.Println("Введите целое положительное число")
	fmt.Scanln(&n)
	fmt.Println(fibo(n))
	fiboMap := map[int]int{}
	for i := 0; i <= n; i++ {
		fiboMap[i] = fibo(i)
	}
	fmt.Println(fiboMap)
}
