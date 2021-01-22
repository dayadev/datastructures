package main

import (
	"fmt"
)

type node struct {
	value string
	left  *node
	right *node
}

func insert(n *node, v string) *node {
	if n == nil {
		return &node{v, nil, nil}
	} else if v <= n.value {
		n.left = insert(n.left, v)
	} else {
		n.right = insert(n.right, v)
	}
	return n
}

/* pre-oder DFS traversal */
func preorder(n *node) {
	if n != nil {
		fmt.Printf(n.value + " ")
		preorder(n.left)
		preorder(n.right)
	}
}

/* in-oder DFS traversal */
func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Printf(n.value + " ")
		inorder(n.right)
	}
}

/* post-oder DFS traversal */
func postorder(n *node) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Printf(n.value + " ")
	}
}

/* breadth first traversal */
func breadth(n *node) {
	if n != nil {
		s := []*node{n}
		for len(s) > 0 {
			current_node := s[0]
			fmt.Printf(current_node.value + " ")
			s = s[1:]
			if current_node.left != nil {
				s = append(s, current_node.left)
			}
			if current_node.right != nil {
				s = append(s, current_node.right)
			}
		}
	}
}
func findMin(n *node) string {
	if n == nil {
		return "0"
	}
	for {
		if n.left == nil {
			return n.value
		}
		n = n.left
	}
}
func findMax(n *node) string {
	if n == nil {
		return "0"
	}
	for {
		if n.right == nil {
			return n.value
		}
		n = n.right
	}
}
func find(n *node, v string) bool {
	if n == nil {
		return false
	}
	if v == n.value {
		return true
	}
	if v < n.value {
		return find(n.left, v)
	}
	if v > n.value {
		return find(n.right, v)
	}
	return false
}
func main() {
	var root *node
	root = insert(root, "F")
	root = insert(root, "D")
	root = insert(root, "H")
	root = insert(root, "B")
	root = insert(root, "E")
	root = insert(root, "G")
	root = insert(root, "J")
	root = insert(root, "Z")
	root = insert(root, "A")

	fmt.Println("Pre-order DFS traversal")
	preorder(root)
	fmt.Println("\nIn-order DFS traversal")
	inorder(root)
	fmt.Println("\nPost-order DFS traversal")
	postorder(root)
	fmt.Println("\nBFS traversal")
	breadth(root)
	fmt.Println("\nMin")
	println(findMin(root))
	fmt.Println("Max")
	println(findMax(root))
	fmt.Println("Existence")
	println(find(root, "X"))
	println(find(root, "Z"))

}
