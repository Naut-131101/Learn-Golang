package main

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList[T]) Search(value T) bool {
	current := l.head

	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}

	return false
}

func (l *LinkedList[T]) Delete(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.data == value {
		l.head = l.head.next

		if l.head == nil {
			l.tail = nil
		}
		return true
	}

	prev := l.head
	current := l.head.next

	for current != nil {
		if current.data == value {
			prev.next = current.next

			if current == l.tail {
				l.tail = prev
			}
			return true
		}

		prev = current
		current = current.next
	}

	return false
}

func (l *LinkedList[T]) PrintList() {
	if l.head == nil {
		fmt.Println("LinkedList is empty")
		return
	}

	current := l.head
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Length() int {
	count := 0
	current := l.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (ll *LinkedList[T]) InsertAt(position int, value T) {
	newNode := &Node[T]{data: value}

	if position <= 0 || ll.head == nil {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	index := 0

	for current.next != nil && index < position-1 {
		current = current.next
		index++
	}

	newNode.next = current.next
	current.next = newNode
}

func (l *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := l.head

	l.tail = l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

func main() {
	list := LinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Prepend(5)

	fmt.Print("List: ")
	list.PrintList()
	
	fmt.Println("Length:", list.Length())

	fmt.Println("Search 20:", list.Search(20))
	fmt.Println("Search 99:", list.Search(99))

	fmt.Println("Delete 20:", list.Delete(20))
	fmt.Print("List after delete 20: ")
	list.PrintList()

	list.InsertAt(1, 15)
	fmt.Print("List after InsertAt(1, 15): ")
	list.PrintList()

	fmt.Print("Before reverse: ")
	list.PrintList()
	list.Reverse()
	fmt.Print("After reverse: ")
	list.PrintList()
}
