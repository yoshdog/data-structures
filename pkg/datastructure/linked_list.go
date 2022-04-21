package datastructure

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Prepend(data T) {
	oldHead := l.head
	l.head = &Node[T]{data: data, next: oldHead}
}

func (l *LinkedList[T]) Delete(data T) {
	currentNode := l.head

	if currentNode.data == data {
		l.head = l.head.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList[T]) Search(data T) bool {
	return false
}
