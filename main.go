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

// func (tree *BinarySearchTree) insertNode(root *Node, new_node *Node) {
// 	if new_node.value < root.value {
// 		if root.left == nil {
// 			root.left = new_node
// 		} else {
// 			tree.insertNode(root.left, new_node)
// 		}
// 	} else {
// 		if root.right == nil {
// 			root.right = new_node
// 		} else {
// 			tree.insertNode(root.right, new_node)
// 		}
// 	}
// }

// func (tree *BinarySearchTree) insert(value int) {
// 	newNode := &Node{value: value}

// 	if tree.root == nil {
// 		tree.root = newNode
// 	} else {
// 		tree.insertNode(tree.root, newNode)
// 	}
// }

// func (tree *BinarySearchTree) searchNode(root *Node, value int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if value < root.value {
// 		return tree.searchNode(root.left, value)
// 	} else if value > root.value {
// 		return tree.searchNode(root.right, value)
// 	} else {
// 		return false
// 	}
// }

// func (tree *BinarySearchTree) search(value int) bool {
// 	return tree.searchNode(tree.root, value)
// }

// func BinarySearchTreeFromArray(array []int) *BinarySearchTree {
// 	bst := BinarySearchTree{}

// 	for _, value := range array {
// 		bst.insert(value)
// 	}

// 	return &bst
// }

func (tree *BinarySearchTree) insert(value int) {
	new_node := &Node{value: value}

	if tree.root == nil {
		tree.root = new_node
		return
	}

	current_node := tree.root

	for {
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

type AVLTreeNode struct {
	value  int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

type AVLTree struct {
	root *AVLTreeNode
}

func (tree *AVLTree) insert(value int) {
	new_node := &AVLTreeNode{value: value}

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

func (tree *AVLTree) search(value int) bool {
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

func main() {
	array := generateSortedArray(100)
	x := BinarySearchTreeFromArray(array)
	fmt.Println(array)
	fmt.Println(x.search(10))
}
