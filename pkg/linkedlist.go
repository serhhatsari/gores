package pkg

/*
Singly Linked List Implementation
Author: Serhat SARI
*/

// LinkedList is a singly linked list
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Node represents a node in the linked list
type Node struct {
	value string
	next  *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add adds a new node with the given value to the beginning of the list
func (ll *LinkedList) Add(value string) {
	// Create a new node
	node := &Node{value: value}

	// Handle the case where the list is empty
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else { // Handle the case where the list is not empty
		node.next = ll.head
		ll.head = node
	}
	ll.size++
}

// Remove removes the node with the given value from the list
func (ll *LinkedList) Remove(value string) bool {
	// Handle the case where the list is empty
	if ll.head == nil {
		return false
	}

	// Handle the case where the head is the node to be removed
	if ll.head.value == value {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		ll.size--
		return true
	}

	// Check the rest of the list
	prev := ll.head
	for current := prev.next; current != nil; current = current.next {
		if current.value == value {
			prev.next = current.next
			if current.next == nil {
				ll.tail = prev
			}
			ll.size--
			return true
		}
		prev = current
	}

	return false
}

// Get returns the value of the node with the given value, if it exists
func (ll *LinkedList) Get(value string) (string, bool) {
	for node := ll.head; node != nil; node = node.next {
		if node.value == value {
			return node.value, true
		}
	}
	return "", false
}

// GetFirst returns the value of the first node, if it exists
func (ll *LinkedList) GetFirst() (string, bool) {
	if ll.head == nil {
		return "", false
	}
	return ll.head.value, true
}

// GetLast returns the value of the last node, if it exists
func (ll *LinkedList) GetLast() (string, bool) {
	if ll.tail == nil {
		return "", false
	}
	return ll.tail.value, true
}

// Size returns the size of the list
func (ll *LinkedList) Size() int {
	return ll.size
}

// Print prints the list
func (ll *LinkedList) Print() {
	for node := ll.head; node != nil; node = node.next {
		println(node.value)
	}
}
