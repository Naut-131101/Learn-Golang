package main

import "fmt"

func main() {
	/**
	Bài 1: Viết chương trình quản lý quán cafe (Coffee Shop Management System) bằng ngôn ngữ Go với yêu cầu sau:
		● Chỉ sử dụng Slice
		● Cho phép thêm mới, xóa, sửa và hiển thị menu
	*/
	var names []string
	var prices []float64

	names = append(names, "Cà phê đen")
	prices = append(prices, 25000)

	names = append(names, "Cà phê sữa")
	prices = append(prices, 29000)

	names = append(names, "Trà đào")
	prices = append(prices, 35000)

	fmt.Println("Menu:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
	}

	editIndex := 1
	names[editIndex] = "Cà phê sữa đá"
	prices[editIndex] = 30000

	fmt.Println("Menu sau khi sửa:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
	}

	deleteIndex := 0
	fmt.Println("Xóa món:", names[deleteIndex])

	names = append(names[:deleteIndex], names[deleteIndex+1:]...)
	prices = append(prices[:deleteIndex], prices[deleteIndex+1:]...)

	fmt.Println("Menu:")
	if len(names) == 0 {
		fmt.Println("  Menu đang trống!")
	} else {
		for i := 0; i < len(names); i++ {
			fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
		}
	}
}
