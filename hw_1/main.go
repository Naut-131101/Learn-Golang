package main

import "fmt"

func main() {
	/**
	Bài tập rèn luyện sử dụng Array: Viết chương trình quản lý quán cafe (Coffee Shop Management System) bằng ngôn ngữ Go với yêu cầu sau:
		1. Chỉ sử dụng Array (Không sử dụng bất kỳ data structure nào khác)
		2. Cho phép thêm mới item, xóa item, sửa item và hiển thị toàn bộ item trong menu
	*/
	var names [10]string
	var prices [10]float64
	count := 0

	names[count] = "Cà phê đen"
	prices[count] = 25000
	count++

	names[count] = "Cà phê sữa"
	prices[count] = 29000
	count++

	names[count] = "Trà đào"
	prices[count] = 35000
	count++

	fmt.Println("Menu:")
	for i := 0; i < count; i++ {
		fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
	}

	editIndex := 1
	names[editIndex] = "Cà phê sữa đá"
	prices[editIndex] = 30000

	fmt.Println("Menu sau khi sửa:")
	for i := 0; i < count; i++ {
		fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
	}

	deleteIndex := 0
	fmt.Println("Xóa món:", names[deleteIndex])

	for i := deleteIndex; i < count-1; i++ {
		names[i] = names[i+1]
		prices[i] = prices[i+1]
	}
	names[count-1] = ""
	prices[count-1] = 0
	count--

	fmt.Println("Menu:")
	if count == 0 {
		fmt.Println("  Menu đang trống!")
	} else {
		for i := 0; i < count; i++ {
			fmt.Printf("  %d. %s - %.0f VND\n", i+1, names[i], prices[i])
		}
	}
}
