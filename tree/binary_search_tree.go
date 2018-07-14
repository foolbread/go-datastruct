package tree

import (
	"fmt"
)

type BinarySearchNode struct {
	left   *BinarySearchNode
	right  *BinarySearchNode
	parent *BinarySearchNode
	data   int
}

func NewBinarySearchNode() *BinarySearchNode {
	r := new(BinarySearchNode)

	return r
}

func (t *BinarySearchNode) SetData(d int) {
	t.data = d
}

func (t *BinarySearchNode) SetLeft(l *BinarySearchNode) {
	t.left = l
}

func (t *BinarySearchNode) SetRight(r *BinarySearchNode) {
	t.right = r
}

func (t *BinarySearchNode) SetParent(p *BinarySearchNode) {
	t.parent = p
}

func (t *BinarySearchNode) GetData() int {
	return t.data
}

func (t *BinarySearchNode) GetLeft() *BinarySearchNode {
	return t.left
}

func (t *BinarySearchNode) GetRight() *BinarySearchNode {
	return t.right
}

func (t *BinarySearchNode) GetParent() *BinarySearchNode {
	return t.parent
}

type BinarySearchTree struct {
	Root *BinarySearchNode
}

func NewBinarySearchTree() *BinarySearchTree {
	r := new(BinarySearchTree)

	return r
}

func (t *BinarySearchTree) InsertData(d int) {
	if t.Root == nil {
		t.Root = NewBinarySearchNode()
		t.Root.SetData(d)
		return
	}

	pos := t.Root
	for {
		if d > pos.GetData() {
			if pos.GetRight() == nil {
				t := NewBinarySearchNode()
				t.SetData(d)
				pos.SetRight(t)
				t.SetParent(pos)
				return
			} else {
				pos = pos.GetRight()
			}
		} else {
			if pos.GetLeft() == nil {
				t := NewBinarySearchNode()
				t.SetData(d)
				pos.SetLeft(t)
				t.SetParent(pos)
				return
			} else {
				pos = pos.GetLeft()
			}
		}
	}

}

func (t *BinarySearchTree) FindData(d int) *BinarySearchNode {
	pos := t.Root
	for pos != nil {
		if pos.GetData() < d {
			pos = pos.GetRight()
		} else if pos.GetData() > d {
			pos = pos.GetLeft()
		} else {
			return pos
		}
	}

	return nil
}

func (t *BinarySearchTree) Min() int {
	if t.Root == nil {
		return -1
	}

	pos := t.Root
	for {
		if pos.GetLeft() == nil {
			return pos.GetData()
		}

		pos = pos.GetLeft()
	}
}

func (t *BinarySearchTree) Max() int {
	if t.Root == nil {
		return -1
	}

	pos := t.Root
	for {
		if pos.GetRight() == nil {
			return pos.GetData()
		}

		pos = pos.GetRight()
	}
}

func (t *BinarySearchTree) Successor(n *BinarySearchNode) *BinarySearchNode {
	if n.GetRight() != nil {
		r := n.GetRight()
		for r.GetLeft() != nil {
			r = r.GetLeft()
		}

		return r
	}

	y := n.GetParent()
	for y != nil && n != y.GetLeft() {
		n = y
		y = n.GetParent()
	}

	return y
}

func (t *BinarySearchTree) Predecessor(n *BinarySearchNode) *BinarySearchNode {
	if n.GetLeft() != nil {
		r := n.GetLeft()
		for r.GetRight() != nil {
			r = r.GetRight()
		}
		return r
	}

	y := n.GetParent()
	for y != nil && n != y.GetRight() {
		n = y
		y = n.GetParent()
	}

	return y
}

func (t *BinarySearchTree) Del(d int) {
	z := t.FindData(d)
	if z == nil {
		return
	}

	if z.GetLeft() == nil {
		t.transplant(z, z.GetRight())
	} else if z.GetRight() == nil {
		t.transplant(z, z.GetLeft())
	} else {
		y := z.GetLeft()
		for y.GetRight() != nil {
			y = y.GetRight()
		}

		if y != z.GetLeft() {
			t.transplant(y, y.GetLeft())
			y.SetLeft(z.GetLeft())
			z.GetLeft().SetParent(y)
		}

		t.transplant(z, y)
		y.SetRight(z.GetRight())
		z.GetRight().SetParent(y)
		/*y := z.GetRight()
		for y.GetLeft() != nil {
			y = y.GetLeft()
		}

		if y != z.GetRight() {
			t.transplant(y, y.GetRight())
			y.SetRight(z.GetRight())
			z.GetRight().SetParent(y)
		}

		t.transplant(z, y)
		y.SetLeft(z.GetLeft())
		z.GetLeft().SetParent(y)*/
	}
}

func (t *BinarySearchTree) transplant(u *BinarySearchNode, v *BinarySearchNode) {
	if u.GetParent() == nil {
		t.Root = v
	} else if u.GetParent().GetLeft() == u {
		u.GetParent().SetLeft(v)
	} else {
		u.GetParent().SetRight(v)
	}

	if v != nil {
		v.SetParent(u.GetParent())
	}
}

func (t *BinarySearchTree) Inorder() {
	t.inorderhelper(t.Root)

	fmt.Println()
}

func (t *BinarySearchTree) inorderhelper(n *BinarySearchNode) {
	if n == nil {
		return
	}

	t.inorderhelper(n.GetLeft())

	fmt.Print(n.GetData(), " ")

	t.inorderhelper(n.GetRight())
}
