package singlelist

import (
	"fmt"
)

func PrintSingleList(l *SingleList) {
	fmt.Print("cur data: ")
	head := l.ListHead()
	for head != nil {
		fmt.Print(head.GetData(), " ")
		head = head.GetNext()
	}
	fmt.Println()
	fmt.Println("singlelist size:", l.GetListSize())
}

type ListElmt struct {
	data int
	next *ListElmt
}

func NewListElmt(d int) *ListElmt {
	r := new(ListElmt)
	r.data = d
	r.next = nil

	return r
}

func (s *ListElmt) GetData() int {
	return s.data
}

func (s *ListElmt) GetNext() *ListElmt {
	return s.next
}

func (s *ListElmt) SetData(d int) {
	s.data = d
}

func (s *ListElmt) SetNext(n *ListElmt) {
	s.next = n
}

type SingleList struct {
	head *ListElmt
	tail *ListElmt
	size int
}

func NewSingleList() *SingleList {
	r := new(SingleList)
	r.head = nil
	r.tail = nil
	r.size = 0

	return r
}

func (s *SingleList) ListHead() *ListElmt {
	return s.head
}

func (s *SingleList) ListTail() *ListElmt {
	return s.tail
}

func (s *SingleList) ListInsert(e *ListElmt) {
	if e == nil {
		return
	}

	s.size++
	if s.tail == nil {
		s.head = e
		s.tail = e
		return
	}

	s.tail.SetNext(e)
	s.tail = e
}

func (s *SingleList) ListDelNext(e *ListElmt) *ListElmt {
	var del *ListElmt
	if e == nil { //delete head
		del = s.head
		if s.head != nil {
			s.head = s.head.GetNext()
			if s.size == 1 {
				s.tail = nil
			}
			s.size--
		}
	} else {
		del = e.GetNext()
		if del != nil {
			e.SetNext(del.GetNext())
			s.size--
		}
	}

	return del
}

func (s *SingleList) GetListSize() int {
	return s.size
}
