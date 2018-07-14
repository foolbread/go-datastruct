// go-program project main.go
package main

import (
	"fmt"
	"github.com/foolbread/go-datastruct/doublelist"
	"github.com/foolbread/go-datastruct/hashtable"
	"github.com/foolbread/go-datastruct/queue"
	"github.com/foolbread/go-datastruct/set"
	"github.com/foolbread/go-datastruct/singlelist"
	"github.com/foolbread/go-datastruct/singlelooplist"
	"github.com/foolbread/go-datastruct/stack"
	"github.com/foolbread/go-datastruct/tree"
	//"math/rand"
)

func main() {
	TestDatastruct()
}

//////////////////////////////////
func TestDatastruct() {
	//TestSingleList()

	//TestDoubleList()

	//TestSingleLoopList()

	//TestStack()

	//TestQueue()

	//TestSet()

	//TestChainHash()

	//TestOpenHash()

	TestBinaryTree()

	//TestBinarySearchTree()
}

func TestAlgoritbms() {

}

/////////////////////////////////
func TestSingleList() {
	fmt.Println("SingleList Test..........")
	l := singlelist.NewSingleList()

	for i := 0; i < 5; i++ {
		l.ListInsert(singlelist.NewListElmt(i))
	}

	singlelist.PrintSingleList(l)

	head := l.ListHead()
	l.ListDelNext(head)

	fmt.Println("after delete....")
	singlelist.PrintSingleList(l)

	l.ListDelNext(nil)
	fmt.Println("after delete head...")
	singlelist.PrintSingleList(l)
}

func TestDoubleList() {
	fmt.Println("DoubleList Test..........")
	l := doublelist.NewDoubleList()

	for i := 0; i < 5; i++ {
		l.ListInsert(doublelist.NewListElmt(i))
	}
	doublelist.PrintDoubleList(l)

	p := l.ListHead().GetNext().GetNext()
	d := doublelist.NewListElmt(7)

	l.ListInsertEx(p, d)
	fmt.Println("after insert data....")
	doublelist.PrintDoubleList(l)

	l.ListDel(d)
	fmt.Println("after delete....")
	doublelist.PrintDoubleList(l)
}

func TestSingleLoopList() {
	fmt.Println("SingleLoopList Test..........")
	l := singlelooplist.NewSingleLoopList()

	head := l.ListHead()
	for i := 0; i < 5; i++ {
		d := singlelooplist.NewListElmt(i)
		l.ListInsertNext(head, d)
		head = d
	}
	singlelooplist.PrintSingleLoopList(l)

	i := l.ListHead().GetNext().GetNext()
	fmt.Println("after insert data....")
	l.ListInsertNext(i, singlelooplist.NewListElmt(17))
	singlelooplist.PrintSingleLoopList(l)

	l.ListDelNext(i)
	fmt.Println("after delete data....")
	singlelooplist.PrintSingleLoopList(l)
}

func TestStack() {
	fmt.Println("Stack Test..........")
	s := stack.NewStack()

	for i := 0; i < 8; i++ {
		s.StackPush(i)
	}

	d, err := s.StackTop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("stack top:", d)
	fmt.Println("stack size:", s.StackSize())

	s.StackPop()

	fmt.Println("after pop data....")

	d, err = s.StackTop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("stack top:", d)
	fmt.Println("stack size:", s.StackSize())
}

func TestQueue() {
	fmt.Println("Queue Test..........")
	q := queue.NewQueue()

	for i := 0; i < 10; i++ {
		q.QueueEnqueue(i)
	}

	h, err := q.QueueHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("queue head:", h)
	fmt.Println("queue size:", q.QueueSize())

	q.QueueDequeue()

	h, err = q.QueueHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("after dequeue data....")

	fmt.Println("queue head:", h)
	fmt.Println("queue size:", q.QueueSize())
}

func TestSet() {
	fmt.Println("Set Test..........")
	s := set.NewSet()
	for i := 0; i < 5; i++ {
		s.SetInsert(i)
	}
	set.PrintSet(s)

	sub := set.NewSet()
	for i := 0; i < 3; i++ {
		sub.SetInsert(i)
	}

	if s.IsSetSubset(sub) {
		fmt.Println("subset!")
	}

	sub.SetInsert(3)
	sub.SetInsert(4)
	if s.IsSetEqual(sub) {
		fmt.Println("set equal!")
	}

	se := set.NewSet()
	for i := 3; i < 10; i++ {
		se.SetInsert(i)
	}
	fmt.Println("difference set:")
	r := s.SetDifference(se)
	set.PrintSet(r)

	fmt.Println("union set:")
	r = s.SetUnion(se)
	set.PrintSet(r)

	fmt.Println("intersection set:")
	r = s.SetIntersection(se)
	set.PrintSet(r)

	fmt.Println("after delete....")
	s.SetRemove(2)
	set.PrintSet(s)

	fmt.Println("after insert....")
	s.SetInsert(7)
	set.PrintSet(s)
}

func TestChainHash() {
	fmt.Println("ChainHashTable Test..........")
	h := hashtable.NewChainHashTable(1699)
	for i := 0; i < 100; i++ {
		h.ChtblInsert(i, i)
	}
	fmt.Println("hash table size:", h.GetHashSize())

	k := 10
	v, b := h.ChtblLookup(k)
	if b {
		fmt.Println("key:", k, "value:", v)
	}

	h.ChtblDel(k)

	v, b = h.ChtblLookup(k)
	if !b {
		fmt.Println("key:", k, "already delete!")
	}

	fmt.Println("after delete one key....")
	fmt.Println("hash table size:", h.GetHashSize())
}

func TestOpenHash() {
	fmt.Println("OpenHashTable Test..........")
	h := hashtable.NewOpenHashTable(1699)

	for i := 0; i < 100; i++ {
		h.OhtblInsert(i, i)
	}
	fmt.Println("hash table size:", h.GetHashSize())

	k := 14
	v, b := h.OhtblLookup(k)
	if b {
		fmt.Println("key:", k, "value:", v)
	}

	h.OhtblDel(k)
	v, b = h.OhtblLookup(k)
	if !b {
		fmt.Println("key:", k, "already delete!")
	}
	fmt.Println("after delete one key....")
	fmt.Println("hash table size:", h.GetHashSize())
}

func TestBinaryTree() {
	fmt.Println("BinaryTree Test..........")
	t := tree.NewBinaryTree()
	for i := 0; i < 10; i++ {
		t.TreeInsert(i)
	}
	fmt.Println("binarytree preorder:")
	t.Preorder()

	fmt.Println("binarytree inorder:")
	t.Inorder()

	fmt.Println("binarytree postorder:")
	t.Postorder()

	fmt.Println("binarytree levelorder:")
	t.Levelorder()
}

func TestBinarySearchTree() {
	fmt.Println("BinarySearchTree Test..........")
	t := tree.NewBinarySearchTree()
	var test []int = []int{15, 6, 18, 3, 7, 17, 20, 2, 4, 13, 9}

	for _, v := range test {
		t.InsertData(v)
	}

	fmt.Println(t.FindData(11))

	fmt.Println("Max:", t.Max())
	fmt.Println("Min:", t.Min())
	t.Inorder()

	fmt.Println("root:", t.Root.GetData())
	fmt.Println("predecessor:", t.Predecessor(t.Root).GetData())
	fmt.Println("successor:", t.Successor(t.Root).GetData())

	t.Del(6)
	t.Inorder()
}
