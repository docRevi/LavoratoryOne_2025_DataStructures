package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) TPUSH(value int) {
	if t.root == nil {
		t.root = &TreeNode{value: value}
		return
	}

	iter := t.root
	var parent *TreeNode

	for iter != nil {
		parent = iter
		if value < iter.value {
			iter = iter.left
		} else if value > iter.value {
			iter = iter.right
		} else {
			return
		}
	}

	if value < parent.value {
		parent.left = &TreeNode{value: value}
	} else {
		parent.right = &TreeNode{value: value}
	}
}

func (t *Tree) TSEARCH(root *TreeNode, value int) bool {
	iter := root
	for iter != nil {
		if iter.value == value {
			return true
		} else if value < iter.value {
			iter = iter.left
		} else {
			iter = iter.right
		}
	}
	return false
}

func (t *Tree) isFullBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.left == nil && root.right == nil {
		return true
	}

	if root.left != nil && root.right != nil {
		return t.isFullBinaryTree(root.left) && t.isFullBinaryTree(root.right)
	}

	return false
}

func (t *Tree) heightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := t.heightTree(root.left)
	rh := t.heightTree(root.right)
	if lh > rh {
		return 1 + lh
	}
	return 1 + rh
}

func (t *Tree) printLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Print(root.value, " ")
	} else {
		t.printLevel(root.left, level-1)
		t.printLevel(root.right, level-1)
	}
}

func (t *Tree) print(root *TreeNode) {
	h := t.heightTree(root)
	for i := 1; i <= h; i++ {
		t.printLevel(root, i)
	}
	fmt.Println()
}

func (t *Tree) printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	t.printInOrder(root.left)
	fmt.Print(root.value, " ")
	t.printInOrder(root.right)
	fmt.Println()
}

func (t *Tree) printPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value, " ")
	t.printPreOrder(root.left)
	t.printPreOrder(root.right)
	fmt.Println()
}

func (t *Tree) printPostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	t.printPostOrder(root.left)
	t.printPostOrder(root.right)
	fmt.Print(root.value, " ")
	fmt.Println()
}
