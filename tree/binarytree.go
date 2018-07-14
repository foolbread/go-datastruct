package tree

import (
	"fmt"
)

type BinaryNode struct {
	data  int
	left  *BinaryNode
	right *BinaryNode
}

func NewBinaryNode(d int, l *BinaryNode, r *BinaryNode) *BinaryNode {
	ret := new(BinaryNode)
	ret.data = d
	ret.left = l
	ret.right = r

	return ret
}

func (t *BinaryNode) SetData(d int) {
	t.data = d
}

func (t *BinaryNode) SetLeft(l *BinaryNode) {
	t.left = l
}

func (t *BinaryNode) SetRight(r *BinaryNode) {
	t.right = r
}

func (t *BinaryNode) GetData() int {
	return t.data
}

func (t *BinaryNode) GetLeft() *BinaryNode {
	return t.left
}

func (t *BinaryNode) GetRight() *BinaryNode {
	return t.right
}

type BinaryTree struct {
	root *BinaryNode
	tree []*BinaryNode
}

func NewBinaryTree() *BinaryTree {
	r := new(BinaryTree)

	return r
}

func (t *BinaryTree) TreeInsert(d int) {
	node := NewBinaryNode(d, nil, nil)
	if t.root == nil {
		t.root = node
	} else {
		l := len(t.tree)
		if l%2 == 0 {
			t.tree[l/2-1].SetRight(node)
		} else {
			t.tree[l/2].SetLeft(node)
		}
	}

	t.tree = append(t.tree, node)
}

func (t *BinaryTree) TreeSize() int {
	return len(t.tree)
}

func (t *BinaryTree) GetTreeRoot() *BinaryNode {
	return t.root
}

func (t *BinaryTree) Preorder() {
	t.preorderhelper(t.root)
	fmt.Println()
}

func (t *BinaryTree) preorderhelper(n *BinaryNode) {
	if n == nil {
		return
	}

	fmt.Print(n.GetData(), " ")

	if n.GetLeft() != nil {
		t.preorderhelper(n.GetLeft())
	}

	if n.GetRight() != nil {
		t.preorderhelper(n.GetRight())
	}

}

func (t *BinaryTree) Inorder() {
	t.inorderhelper(t.root)
	fmt.Println()
}

func (t *BinaryTree) inorderhelper(n *BinaryNode) {
	if n == nil {
		return
	}

	if n.GetLeft() != nil {
		t.inorderhelper(n.GetLeft())
	}

	fmt.Print(n.GetData(), " ")

	if n.GetRight() != nil {
		t.inorderhelper(n.GetRight())
	}
}

func (t *BinaryTree) Postorder() {
	t.postorderhelper(t.root)
	fmt.Println()
}

func (t *BinaryTree) postorderhelper(n *BinaryNode) {
	if n == nil {
		return
	}

	if n.GetLeft() != nil {
		t.postorderhelper(n.GetLeft())
	}

	if n.GetRight() != nil {
		t.postorderhelper(n.GetRight())
	}

	fmt.Print(n.GetData(), " ")
}

func (t *BinaryTree) Levelorder() {
	for _, v := range t.tree {
		fmt.Print(v.GetData(), " ")
	}
	fmt.Println()
}
