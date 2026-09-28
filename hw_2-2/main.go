package main

import "fmt"

func main() {
	/**
	Bài 2: Viết chương trình quản lý quán cafe (Coffee Shop Management System) bằng ngôn ngữ Go với yêu cầu sau:
		● Chỉ sử dụng Map
		● Cho phép thêm mới, xóa, sửa và hiển thị menu
	*/
	menu := make(map[string]float64)

	menu["Cà phê đen"] = 25000
	menu["Cà phê sữa"] = 29000
	menu["Trà đào"] = 35000

	fmt.Println("Menu sau khi thêm:")
	for name, price := range menu {
		fmt.Printf("  - %s: %.0f VND\n", name, price)
	}

	menu["Cà phê sữa"] = 30000

	oldPrice, ok := menu["Trà đào"]
	if ok {
		menu["Trà đào cam sả"] = oldPrice
		delete(menu, "Trà đào")
	}
	fmt.Println("Menu sau khi sửa:")
	for name, price := range menu {
		fmt.Printf("  - %s: %.0f VND\n", name, price)
	}

	delete(menu, "Cà phê đen")

	fmt.Println("Menu:")
	if len(menu) == 0 {
		fmt.Println("  Menu đang trống!")
	} else {
		for name, price := range menu {
			fmt.Printf("  - %s: %.0f VND\n", name, price)
		}
	}
}
