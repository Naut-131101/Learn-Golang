package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) Search(value int) bool {
	current := l.head

	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}

	return false
}

func (l *LinkedList) Delete(value int) bool {
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

func (l *LinkedList) PrintList() {
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

// Length: Đếm số node
func (l *LinkedList) Length() int {
	count := 0
	current := l.head

	for current != nil {
		count++
		current = current.next
	}

	return count
}

// InsertAt: Chèn value vào vị trí position bất kỳ
func (ll *LinkedList) InsertAt(position int, value int) {
	newNode := &Node{data: value}

	if position <= 0 || ll.head == nil {
		// Chèn đầu danh sách
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	index := 0

	// Duyệt đến node ngay trước vị trí cần chèn
	for current.next != nil && index < position-1 {
		current = current.next
		index++
	}

	newNode.next = current.next
	current.next = newNode
}

func main() {
	list := LinkedList{}

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

	fmt.Println("Delete 5:", list.Delete(5))
	fmt.Print("List after delete 5: ")
	list.PrintList()

	fmt.Println("Delete 30:", list.Delete(30))
	fmt.Print("List after delete 30: ")
	list.PrintList()

	fmt.Println("Length:", list.Length())
}
