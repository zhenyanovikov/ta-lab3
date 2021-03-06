package ta_lab3

import "fmt"

type BiDirectionalNode struct {
	el   interface{}
	next *BiDirectionalNode
	prev *BiDirectionalNode
}

type DoubleLinkedList struct {
	firstSentinel *BiDirectionalNode
	last          *BiDirectionalNode
	length        int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	sentinelNode := &BiDirectionalNode{}
	return &DoubleLinkedList{sentinelNode, sentinelNode, 0}
}

func NewBiDirectionalNode(el interface{}) *BiDirectionalNode {
	return &BiDirectionalNode{
		el:   el,
		next: nil,
		prev: nil,
	}
}

func (l *DoubleLinkedList) Add(el interface{}) {
	newNode := NewBiDirectionalNode(el)

	lastNode := l.Last()
	lastNode.next = newNode
	newNode.prev = lastNode

	l.last = newNode
	l.length++
}

func (l *DoubleLinkedList) Insert(index int, el interface{}) {
	if index >= l.Length() {
		l.Add(el)
	} else {
		node := l.getNode(index)
		newNode := NewBiDirectionalNode(el)

		node.prev.next = newNode
		newNode.prev = node.prev
		node.prev = newNode
		newNode.next = node

		l.length++
	}

}

func (l *DoubleLinkedList) Get(index int) interface{} {
	return l.getNode(index).el
}

func (l *DoubleLinkedList) getNode(index int) *BiDirectionalNode {
	if index >= l.Length() || index < 0 {
		fmt.Println("index = ! ", index, "| ! |  ", l.length)
		panic("index out of bounds")
	}

	node := l.first()
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *DoubleLinkedList) Find(el interface{}) (index int, ok bool) {
	node := l.firstSentinel

	for i := 0; i < l.Length(); i++ {
		node = node.next
		if node.el == el {
			return i, true
		}
	}
	return -1, false
}

func (l *DoubleLinkedList) Remove(index int) {
	node := l.getNode(index)
	if index == l.Length()-1 {
		l.last = node.prev
	} else {
		node.next.prev = node.prev
	}

	node.prev.next = node.next
	l.length--
}

func (l *DoubleLinkedList) Length() int {
	return l.length
}

func (l *DoubleLinkedList) first() *BiDirectionalNode {
	if l.firstSentinel.next == nil {
		return l.firstSentinel
	} else {
		return l.firstSentinel.next
	}

}

func (l *DoubleLinkedList) Last() *BiDirectionalNode {
	return l.last
}

func (l *DoubleLinkedList) Copy() List {
	list := NewDoubleLinkedList()
	for i := 0; i < l.length; i++ {
		list.Add(l.Get(i))
	}
	return list
}
