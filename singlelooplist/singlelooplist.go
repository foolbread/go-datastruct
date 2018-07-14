package singlelooplist

import (
	"fmt"
)

func PrintSingleLoopList(l *SingleLoopList) {
	fmt.Print("cur data: ")
	s := l.GetSize()
	head := l.ListHead()
	for i := 0; i < s; i++ {
		fmt.Print(head.GetData(), " ")
		head = head.GetNext()
	}
	fmt.Println()
	fmt.Println("singlelooplist size:", l.GetSize())
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

func (s *ListElmt) SetNext(e *ListElmt) {
	s.next = e
}

type SingleLoopList struct {
	head *ListElmt
	size int
}

func NewSingleLoopList() *SingleLoopList {
	r := new(SingleLoopList)
	r.head = nil
	r.size = 0

	return r
}

func (s *SingleLoopList) GetSize() int {
	return s.size
}

func (s *SingleLoopList) ListHead() *ListElmt {
	return s.head
}

func (s *SingleLoopList) ListDelNext(e *ListElmt) *ListElmt {
	if e == nil || s.size == 0 {
		return nil
	}

	var del *ListElmt
	if e.GetNext() == e {
		del = e
		s.head = nil
	} else {
		del = e.GetNext()
		if del != nil {
			e.SetNext(del.GetNext())
		}
	}
	s.size--

	return del
}

func (s *SingleLoopList) ListInsertNext(e *ListElmt, i *ListElmt) {
	if i == nil {
		return
	}

	if s.size == 0 {
		s.head = i
		i.SetNext(i)
	} else {
		i.SetNext(e.GetNext())
		e.SetNext(i)
	}

	s.size++
	return
}
