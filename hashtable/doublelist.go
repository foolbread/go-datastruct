package hashtable

type ListElmt struct {
	key  int
	data int
	pre  *ListElmt
	next *ListElmt
}

func NewListElmt(k int, d int) *ListElmt {
	r := new(ListElmt)
	r.key = k
	r.data = d
	r.pre = nil
	r.next = nil

	return r
}

func (d *ListElmt) GetKey() int {
	return d.key
}

func (d *ListElmt) GetData() int {
	return d.data
}

func (d *ListElmt) GetPre() *ListElmt {
	return d.pre
}

func (d *ListElmt) GetNext() *ListElmt {
	return d.next
}

func (d *ListElmt) SetKey(k int) {
	d.key = k
}

func (d *ListElmt) SetData(da int) {
	d.data = da
}

func (d *ListElmt) SetPre(e *ListElmt) {
	d.pre = e
}

func (d *ListElmt) SetNext(e *ListElmt) {
	d.next = e
}

type DoubleList struct {
	head *ListElmt
	tail *ListElmt
	size int
}

func NewDoubleList() *DoubleList {
	r := new(DoubleList)
	r.head = nil
	r.tail = nil
	r.size = 0

	return r
}

func (d *DoubleList) GetListSize() int {
	return d.size
}

func (d *DoubleList) ListHead() *ListElmt {
	return d.head
}

func (d *DoubleList) ListTail() *ListElmt {
	return d.tail
}

func (d *DoubleList) ListInsertEx(e *ListElmt, i *ListElmt) {
	if e == nil {
		return
	}

	next := e.GetNext()

	e.SetNext(i)
	i.SetPre(e)

	i.SetNext(next)
	if next != nil {
		next.SetPre(i)
	} else {
		d.tail = i
	}

	d.size++
}

func (d *DoubleList) ListInsert(e *ListElmt) {
	if e == nil {
		return
	}

	if d.head == nil {
		d.head = e
		d.tail = e
		d.size++
		return
	}

	e.SetPre(d.tail)
	d.tail.SetNext(e)
	d.tail = e
	d.size++
}

func (d *DoubleList) ListDel(e *ListElmt) {
	if e == nil || d.size == 0 {
		return
	}

	if e == d.head {
		d.head = d.head.GetNext()
		if d.head == nil {
			d.tail = nil
		} else {
			d.head.SetPre(nil)
		}
	} else {
		e.GetPre().SetNext(e.GetNext())
		if e.GetNext() != nil {
			e.GetNext().SetPre(e.GetPre())
		} else {
			d.tail = e.GetPre()
		}
	}

	d.size--
}
