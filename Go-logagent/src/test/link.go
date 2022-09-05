package main

import "fmt"

type LinkNode struct {
	next *LinkNode
	data interface{}
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data:data,
		next: nil,
	}
	if p.tail == nil && p.head == nil  {
		p.tail, p.head = node, node
		return
	}

	node.next = p.head
	p.head = node

}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data:data,
		next: nil,
	}
	if p.tail == nil && p.head == nil  {
		p.tail, p.head = node, node
		return
	}

	p.tail.next = node
	p.tail = node

}

func (p *Link) Trans() {
	s := p.head
	for s != nil {
		fmt.Println(s.data)
		s = s.next
	}
}
