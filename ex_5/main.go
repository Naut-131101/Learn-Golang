package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (v+1)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) Walk(ch chan int) {
	var walk func(*Tree)

	walk = func(t *Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}

	walk(t)
	close(ch)
}

func (t1 *Tree) Same(t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go t1.Walk(ch1)
	go t2.Walk(ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(rand.Perm(10))
	fmt.Println("")

	t1 := New(1)
	t2 := New(2)

	ch1 := make(chan int)
	go t1.Walk(ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
	fmt.Println("")

	ch2 := make(chan int)
	go t2.Walk(ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("")

	fmt.Println(t1.Same(t2))
	fmt.Println(t1.Same(t1))
}
