package hashtable

type ChainHashTable struct {
	data []DoubleList
	size int
}

func NewChainHashTable(b int) *ChainHashTable {
	r := new(ChainHashTable)
	r.data = make([]DoubleList, b, b)
	r.size = 0

	return r
}

func (h *ChainHashTable) ChtblInsert(k int, v int) {
	_, be := h.ChtblLookup(k)
	if be {
		return
	}

	idx := h.hashbuckets(k)
	h.data[idx].ListInsert(NewListElmt(k, v))
	h.size++
}

func (h *ChainHashTable) ChtblDel(k int) {
	idx := h.hashbuckets(k)
	head := h.data[idx].ListHead()
	for head != nil {
		if head.GetKey() == k {
			h.data[idx].ListDel(head)
			h.size--
			return
		}
		head = head.GetNext()
	}
}

func (h *ChainHashTable) ChtblLookup(k int) (int, bool) {
	idx := h.hashbuckets(k)
	head := h.data[idx].ListHead()
	for head != nil {
		if head.GetKey() == k {
			return head.GetData(), true
		}
		head = head.GetNext()
	}
	return 0, false
}

func (h *ChainHashTable) GetHashSize() int {
	return h.size
}

func (h *ChainHashTable) hashbuckets(k int) int {
	return k % len(h.data)
}
