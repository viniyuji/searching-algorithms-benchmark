package main

import (
	"fmt"
	"math/rand"
)

func generateSortedArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = i + 1
	}
	return array
}

func generateReversedSortedArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = size - i
	}
	return array
}

func generateRandomSortedArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size) + 1
	}
	return array
}

func generateAlmostSortedArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = i + 1
	}

	for i := 0; i < size/10; i++ {
		random_index := rand.Intn(size)
		random_index_2 := rand.Intn(size)
		array[random_index], array[random_index_2] = array[random_index_2], array[random_index]
	}

	return array
}

func linearSearch(array []int, element int) int {
	for index, value := range array {
		if value == element {
			return index
		}
	}
	return -1
}

func binarySearch(array []int, element int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		middle := (left + right) / 2

		if array[middle] < element {
			left = middle + 1
		} else if array[middle] > element {
			right = middle - 1
		} else {
			return middle
		}
	}

	return -1
}

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) insert(value int) {
	new_node := &Node{value: value}

	if tree.root == nil {
		tree.root = new_node
		return
	}

	current_node := tree.root

	for true {
		if value < current_node.value {
			if current_node.left == nil {
				current_node.left = new_node
				return
			}
			current_node = current_node.left
		} else {
			if current_node.right == nil {
				current_node.right = new_node
				return
			}
			current_node = current_node.right
		}
	}
}

func (tree *BinarySearchTree) search(value int) bool {
	current_node := tree.root

	for current_node != nil {
		if value < current_node.value {
			current_node = current_node.left
		} else if value > current_node.value {
			current_node = current_node.right
		} else {
			return true
		}
	}

	return false
}

func BinarySearchTreeFromArray(array []int) *BinarySearchTree {
	bst := BinarySearchTree{}

	for _, value := range array {
		bst.insert(value)
	}

	return &bst
}

// AVL Tree

type AVLTree struct {
	root *Node
}

func (tree *AVLTree) height(node *Node) int {
	if node == nil {
		return -1
	}

	left_height := tree.height(node.left)
	right_height := tree.height(node.right)

	if left_height > right_height {
		return left_height + 1
	} else {
		return right_height + 1
	}
}

func (tree *AVLTree) balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}

	return tree.height(node.left) - tree.height(node.right)
}

func (tree *AVLTree) rightRotate(node *Node) *Node {
	new_root := node.right
	node.right = new_root.left
	new_root.left = node
	return new_root
}

func (tree *AVLTree) leftRotate(node *Node) *Node {
	new_root := node.left
	node.left = new_root.right
	new_root.right = node
	return new_root
}

func (tree *AVLTree) insertNode(root *Node, value int) *Node {

	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = tree.insertNode(root.left, value)
	} else {
		root.right = tree.insertNode(root.right, value)
	}

	balance_factor := tree.balanceFactor(root)

	if balance_factor > 1 {
		if value < root.left.value {
			return tree.leftRotate(root)
		} else {
			root.left = tree.rightRotate(root.left)
			return tree.leftRotate(root)
		}
	}

	if balance_factor < -1 {
		if value > root.right.value {
			return tree.rightRotate(root)
		} else {
			root.right = tree.leftRotate(root.right)
			return tree.rightRotate(root)
		}
	}

	return root
}

func (tree *AVLTree) insert(value int) {
	tree.root = tree.insertNode(tree.root, value)
}

func (tree *AVLTree) searchNode(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if value < root.value {
		return tree.searchNode(root.left, value)
	} else if value > root.value {
		return tree.searchNode(root.right, value)
	} else {
		return true
	}
}

func (tree *AVLTree) search(value int) bool {
	return tree.searchNode(tree.root, value)
}

func AVLTreeFromArray(array []int) *AVLTree {
	avl := AVLTree{}

	for _, value := range array {
		avl.insert(value)
	}

	return &avl
}

func main() {
	array := generateSortedArray(100)
	x := AVLTreeFromArray(array)
	fmt.Println(array)
	fmt.Println(x.search(101))
}
