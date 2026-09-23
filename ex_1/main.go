package main

import (
	"fmt"
)

func main() {
	// 1. Tạo slice rỗng, append 1,2,3,4,5 rồi in len/cap sau mỗi lần append.
	s := make([]int, 0)
	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))
	}
	fmt.Println("---------------------------------")

	// 2. Cắt []int{10,20,30,40,50} thành [20,30,40].
	t := []int{10, 20, 30, 40, 50}
	tslice := t[1:4]
	fmt.Println("slice:", t)
	fmt.Println("slice da cat:", tslice)
	fmt.Println("---------------------------------")

	// 3. Copy một slice string sang slice mới và chứng minh hai slice có thể thay đổi độc lập.
	s1 := []string{"a", "b", "c"}
	s1copy := append([]string(nil), s1...)
	s1[0] = "Tuan"
	s1copy[1] = "Tuan"
	fmt.Println("s1:", s1, "s1copy:", s1copy)
}
