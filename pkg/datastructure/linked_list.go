package datastructure

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Prepend(value T) {
	oldHead := l.head
	l.head = &Node[T]{value: value, next: oldHead}
}

func (l *LinkedList[T]) Delete(value T) {
	currentNode := l.head

	if currentNode.value == value {
		l.head = l.head.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.value == value {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (l *LinkedList[T]) Search(value T) bool {
	return false
}
