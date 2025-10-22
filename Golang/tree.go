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

var tree Tree

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

func (t *Tree) TSEARCH(value int) bool {
	iter := t.root

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

func (t *Tree) IsFullBinaryTree() bool {
	return isFullBinaryTree(t.root)
}

func isFullBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.left == nil && root.right == nil {
		return true
	}

	if root.left != nil && root.right != nil {
		return isFullBinaryTree(root.left) && isFullBinaryTree(root.right)
	}

	return false
}

func (t *Tree) HeightTree() int {
	return heightTree(t.root)
}

func heightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := heightTree(root.left)
	rh := heightTree(root.right)
	if lh > rh {
		return 1 + lh
	}
	return 1 + rh
}

func (t *Tree) PrintLevel(level int) {
	printLevel(t.root, level)
}

func printLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Print(root.value, " ")
	} else {
		printLevel(root.left, level-1)
		printLevel(root.right, level-1)
	}
}

func (t *Tree) Print() {
	h := t.HeightTree()
	for i := 1; i <= h; i++ {
		t.PrintLevel(i)
	}
	fmt.Println()
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.left)
	fmt.Print(root.value, " ")
	printInOrder(root.right)
}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value, " ")
	printPreOrder(root.left)
	printPreOrder(root.right)
}

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
	fmt.Println()
}

func printPostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printPostOrder(root.left)
	printPostOrder(root.right)
	fmt.Print(root.value, " ")
}
