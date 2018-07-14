package set

import (
	"fmt"
	"go-program/datastruct/doublelist"
)

func PrintSet(s *Set) {
	head := s.data.ListHead()
	fmt.Print("cur data: ")
	for head != nil {
		fmt.Print(head.GetData(), " ")
		head = head.GetNext()
	}
	fmt.Println()

	fmt.Println("set size:", s.GetSetSize())
}

type Set struct {
	data *doublelist.DoubleList
}

func NewSet() *Set {
	r := new(Set)
	r.data = doublelist.NewDoubleList()

	return r
}

func (s *Set) SetUnion(se *Set) *Set {
	if se == nil {
		return s
	}

	r := NewSet()
	head := s.data.ListHead()
	for head != nil {
		r.SetInsert(head.GetData())
		head = head.GetNext()
	}

	head = se.data.ListHead()
	for head != nil {
		if !r.IsSetMember(head.GetData()) {
			r.SetInsert(head.GetData())
		}
		head = head.GetNext()
	}

	return r
}

func (s *Set) SetIntersection(se *Set) *Set {
	if se == nil {
		return nil
	}

	r := NewSet()
	head := s.data.ListHead()
	for head != nil {
		if se.IsSetMember(head.GetData()) {
			r.SetInsert(head.GetData())
		}
		head = head.GetNext()
	}

	return r
}

func (s *Set) SetDifference(se *Set) *Set {
	if se == nil {
		return s
	}

	r := NewSet()
	head := s.data.ListHead()
	for head != nil {
		if !se.IsSetMember(head.GetData()) {
			r.SetInsert(head.GetData())
		}
		head = head.GetNext()
	}

	return r
}

func (s *Set) SetInsert(d int) {
	s.data.ListInsert(doublelist.NewListElmt(d))
}

func (s *Set) SetRemove(d int) {
	head := s.data.ListHead()
	for head != nil {
		if head.GetData() == d {
			s.data.ListDel(head)
			return
		}
		head = head.GetNext()
	}
}

func (s *Set) IsSetMember(d int) bool {
	head := s.data.ListHead()
	for head != nil {
		if head.GetData() == d {
			return true
		}
		head = head.GetNext()
	}

	return false
}

func (s *Set) IsSetSubset(se *Set) bool {
	if se == nil {
		return false
	}

	head := se.data.ListHead()
	for head != nil {
		if !s.IsSetMember(head.GetData()) {
			return false
		}
		head = head.GetNext()
	}

	return true
}

func (s *Set) IsSetEqual(se *Set) bool {
	if se == s {
		return true
	}

	if se.GetSetSize() != s.GetSetSize() {
		return false
	}

	head := se.data.ListHead()
	for head != nil {
		if !s.IsSetMember(head.GetData()) {
			return false
		}
		head = head.GetNext()
	}

	return true
}

func (s *Set) GetSetSize() int {
	return s.data.GetListSize()
}
