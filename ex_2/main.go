package main

import "fmt"

func main() {
	fmt.Println("sumTo:", sumTo(5))
	fmt.Println("---------------------------------")
	countdown(5)
	fmt.Println("---------------------------------")
	fmt.Println("fibonacci:", fibonacci(5))
}

func sumTo(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumTo(n-1)
}

func countdown(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("CountDown:", n)
	return countdown(n - 1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		tong := a + b
		fmt.Printf("Fibo: %d + %d = %d\n", a, b, tong)
		a = b
		b = tong
	}

	return b
}
