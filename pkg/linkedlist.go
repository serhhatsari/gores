package pkg

/*
Doubly Linked List Implementation
Author: Serhat SARI
*/

// LinkedList is a doubly linked list
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Node represents a node in the linked list
type Node struct {
	value string
	prev  *Node
	next  *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Push adds a new node with the given value to the beginning of the list
func (ll *LinkedList) Push(value string) {
	// Create a new node
	node := &Node{value: value}

	// Handle the case where the list is empty
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else { // Handle the case where the list is not empty
		node.next = ll.head
		ll.head.prev = node
		ll.head = node
	}
	ll.size++
}

// PushBack adds a new node with the given value to the end of the list
func (ll *LinkedList) PushBack(value string) {
	// Create a new node
	node := &Node{value: value}

	// Handle the case where the list is empty
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else { // Handle the case where the list is not empty
		ll.tail.next = node
		node.prev = ll.tail
		ll.tail = node
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
	for current := ll.head.next; current != nil; current = current.next {
		if current.value == value {
			ll.head.prev.next = current.next
			if current.next == nil {
				ll.tail = ll.head.prev
			}
			ll.size--
			return true
		}
	}

	return false
}

// RemoveFirst removes the first node from the list
func (ll *LinkedList) RemoveFirst() (string, bool) {
	// Handle the case where the list is empty
	if ll.head == nil {
		return "", false
	}

	// Handle the case where there is only one node in the list
	if ll.size == 1 {
		value := ll.head.value
		ll.head = nil
		ll.tail = nil
		ll.size--
		return value, true
	}

	// Handle the case where there are more than one nodes in the list
	value := ll.head.value
	ll.head = ll.head.next
	ll.size--
	return value, true
}

// RemoveLast removes the last node from the list
func (ll *LinkedList) RemoveLast() (string, bool) {
	// Handle the case where the list is empty
	if ll.tail == nil {
		return "", false
	}

	// Handle the case where there is only one node in the list
	if ll.head == ll.tail {
		value := ll.head.value
		ll.head = nil
		ll.tail = nil
		ll.size--
		return value, true
	}

	// Handle the case where there are more than one nodes in the list
	value := ll.tail.value
	ll.tail.prev.next = nil
	ll.tail = ll.tail.prev
	ll.size--
	return value, true
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

func (ll *LinkedList) Range(start int, end int) []string {
	var result []string
	if start < 0 {
		start = ll.size + start
		if start < 0 {
			start = 0
		}
	}

	if end < 0 {
		end = ll.size + end
		if end < 0 {
			return result
		}
	}

	if start >= ll.size {
		return result
	}

	if end >= ll.size {
		end = ll.size - 1
	}

	if start > end {
		return result
	}

	count := 0
	for node := ll.head; node != nil; node = node.next {
		if count >= start && count <= end {
			result = append(result, node.value)
		}
		count++
	}

	return result
}

func (ll *LinkedList) Trim(start, end int) {
	if start < 0 {
		start = ll.size + start
		if start < 0 {
			start = 0
		}
	}

	if end < 0 {
		end = ll.size + end
		if end < 0 {
			return
		}
	}

	if start >= ll.size {
		return
	}

	if end >= ll.size {
		end = ll.size - 1
	}

	if start > end {
		return
	}

	count := 0
	for node := ll.head; node != nil; {
		if count >= start && count <= end {
			node = node.next
			count++
			continue
		}
		if node == ll.head {
			ll.head = node.next
			node = node.next
			count++
			continue
		}

		node.prev.next = node.next
		if node.next == nil {
			ll.tail = node.prev
		} else {
			node.next.prev = node.prev
		}
		count++
	}

	ll.size = end - start + 1
}

// Print prints the list
func (ll *LinkedList) Print() {
	for node := ll.head; node != nil; node = node.next {
		println(node.value)
	}
}
