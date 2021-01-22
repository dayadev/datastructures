package main

func main() {
	tree := &Tree{}
	keys := []int{10, 5, 15, 3, 7, 18}
	for _, k := range keys {
		tree.Insert(k)
	}
	println(tree.RangeSum(tree.root, 7, 15))
	// println(tree.FindMin())
	// println(tree.FindMax())
	// println(tree.Find(-1))
	// println(tree.Find(0))
	// tree.Inorder(tree.root)
}

//Node is a node of tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

//Tree is a tree
type Tree struct {
	root *Node
}

// Insert inserts into tree
func (t *Tree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{val: v}
	}
	t.root.Insert(v)
}

// Insert inserts into tree
func (n *Node) Insert(key int) *Node {
	if n == nil {
		return &Node{
			val:   key,
			left:  nil,
			right: nil,
		}
	}
	if key < n.val {
		if n.left == nil {
			n.left = &Node{val: key}
		} else {
			return n.left.Insert(key)
		}
	}
	if key > n.val {
		if n.right == nil {
			n.right = &Node{val: key}
		} else {
			return n.right.Insert(key)
		}
	}
	return n
}

// Inorder traverses inorder
func (t *Tree) Inorder(n *Node) {
	if n == nil {
		return
	}
	t.Inorder(n.left)
	println(n.val)
	t.Inorder(n.right)
}

// Postorder traverses postorder
func (t *Tree) Postorder(n *Node) {
	if n == nil {
		return
	}

	t.Postorder(n.left)
	t.Postorder(n.right)
	println(n.val)

}

// Preorder traverses Preorder
func (t *Tree) Preorder(n *Node) {
	if n == nil {
		return
	}
	println(n.val)
	t.Preorder(n.left)
	t.Preorder(n.right)

}
func (n *Node) find(v int) bool {
	if n == nil {
		return false
	}
	if v == n.val {
		return true
	}
	if v < n.val {
		return n.left.find(v)
	}
	if v > n.val {
		return n.right.find(v)
	}
	return false
}

//Find finds a node
func (t *Tree) Find(v int) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(v)
}

//FindMin finds min
func (t *Tree) FindMin() int {
	if t == nil {
		return 0
	}
	n := t.root
	if n == nil {
		return 0
	}
	for {
		if n.left == nil {
			return n.val
		}
		n = n.left
	}
}

//FindMax finds max
func (t *Tree) FindMax() int {
	if t == nil {
		return 0
	}
	n := t.root
	if n == nil {
		return 0
	}
	for {
		if n.right == nil {
			return n.val
		}
		n = n.right
	}
}

// RangeSum return sum for threshold low,high
func (t *Tree) RangeSum(root *Node, low, high int) int {
	if root == nil {
		return 0
	}
	if root.val < low {
		return t.RangeSum(root.right, low, high)
	}
	if root.val > high {
		return t.RangeSum(root.left, low, high)
	}
	return root.val + t.RangeSum(root.left, low, high) + t.RangeSum(root.right, low, high)
}
