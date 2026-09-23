package main

import "fmt"

func main() {
	fmt.Println("sumTo:", sumTo(5))
	fmt.Println("---------------------------------")
	countdown(5)
	fmt.Println("---------------------------------")
	fmt.Println("fibonacci:", fibonacci(5))
}

// 1. Viết recursive function sumTo(n) tính 1+...+n.
func sumTo(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumTo(n-1)
}

// 2. Viết recursive function countdown(n) in từ n về 1.
func countdown(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("CountDown:", n)
	return countdown(n - 1)
}

// 3. Viết 2 bai tap tren bang phiên bản Fibonacci co vòng lặp để so sánh. viet giong vi du ben duoi
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
