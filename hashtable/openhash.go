package hashtable

type valuePair struct {
	key  int
	data int
}

func newValuePair(k int, d int) *valuePair {
	r := new(valuePair)
	r.key = k
	r.data = d

	return r
}

type OpenHashTable struct {
	data  []*valuePair
	csize int
	size  int
}

func NewOpenHashTable(s int) *OpenHashTable {
	r := new(OpenHashTable)
	r.data = make([]*valuePair, s, s)
	r.csize = s
	r.size = 0

	return r
}

func (h *OpenHashTable) OhtblInsert(k int, d int) {
	if h.csize == h.size {
		return
	}

	_, b := h.OhtblLookup(k)
	if b {
		return
	}

	idx := 0
	for i := 0; i < h.csize; i++ {
		idx = h.hashkey(k, i)
		if h.data[idx] == nil {
			h.data[idx] = newValuePair(k, d)
			h.size++
			return
		}
	}
}

func (h *OpenHashTable) OhtblDel(k int) {
	idx := 0
	for i := 0; i < h.csize; i++ {
		idx = h.hashkey(k, i)
		if h.data[idx] != nil && h.data[idx].key == k {
			h.data[idx] = nil
			h.size--
			return
		}
	}
}

func (h *OpenHashTable) OhtblLookup(k int) (int, bool) {
	idx := 0
	for i := 0; i < h.csize; i++ {
		idx = h.hashkey(k, i)
		if h.data[idx] != nil && h.data[idx].key == k {
			return h.data[idx].data, true
		}
	}

	return 0, false
}

func (h *OpenHashTable) GetHashSize() int {
	return h.size
}

func (h *OpenHashTable) hashkey(k int, i int) int {
	return (h.hashkeyhelper1(k) + i*h.hashkeyhelper2(k)) % h.csize
}

func (h *OpenHashTable) hashkeyhelper1(k int) int {
	return k % h.csize
}

func (h *OpenHashTable) hashkeyhelper2(k int) int {
	return 1 + k%(h.csize-1)
}
