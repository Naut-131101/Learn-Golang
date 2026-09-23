package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))
	}
	fmt.Println("---------------------------------")

	t := []int{10, 20, 30, 40, 50}
	tslice := t[1:4]
	fmt.Println("slice:", t)
	fmt.Println("slice da cat:", tslice)
	fmt.Println("---------------------------------")

	s1 := []string{"a", "b", "c"}
	s1copy := append([]string(nil), s1...)
	s1[0] = "Tuan"
	s1copy[1] = "Tuan"
	fmt.Println("s1:", s1, "s1copy:", s1copy)
}
