package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// Array Generators

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

// Linear Search

func LinearSearch(array []int, element int) int {
	for index, value := range array {
		if value == element {
			return index
		}
	}
	return -1
}

// Binary Search

func BinarySearch(array []int, element int) int {
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

// Binary Search Tree

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) Insert(value int) {
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

func (tree *BinarySearchTree) Search(value int) bool {
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
		bst.Insert(value)
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

func (tree *AVLTree) Insert(value int) {
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

func (tree *AVLTree) Search(value int) bool {
	return tree.searchNode(tree.root, value)
}

func AVLTreeFromArray(array []int) *AVLTree {
	avl := AVLTree{}

	for _, value := range array {
		avl.Insert(value)
	}

	return &avl
}

// Red-Black Tree
type Color bool

const (
	Red   Color = false
	Black Color = true
)

type redBlackNode struct {
	value  int
	color  Color
	parent *redBlackNode
	left   *redBlackNode
	right  *redBlackNode
}

type RedBlackTree struct {
	root *redBlackNode
}

func NewNode(value int, color Color) *redBlackNode {
	return &redBlackNode{
		value: value,
		color: color,
	}
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (tree *RedBlackTree) Insert(value int) {
	newNode := NewNode(value, Red)
	if tree.root == nil {
		tree.root = newNode
	} else {
		insertNode(tree.root, newNode)
	}
	tree.fixInsert(newNode)
}

func insertNode(root, node *redBlackNode) {
	if node.value < root.value {
		if root.left == nil {
			root.left = node
			node.parent = root
		} else {
			insertNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
			node.parent = root
		} else {
			insertNode(root.right, node)
		}
	}
}

func (tree *RedBlackTree) fixInsert(node *redBlackNode) {
	for node != tree.root && node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == Red {
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.leftRotate(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.rightRotate(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == Red {
				node.parent.color = Black
				uncle.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rightRotate(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.leftRotate(node.parent.parent)
			}
		}
	}
	tree.root.color = Black
}

func (tree *RedBlackTree) leftRotate(node *redBlackNode) {
	new_node := node.right
	node.right = new_node.left
	if new_node.left != nil {
		new_node.left.parent = node
	}
	new_node.parent = node.parent
	if node.parent == nil {
		tree.root = new_node
	} else if node == node.parent.left {
		node.parent.left = new_node
	} else {
		node.parent.right = new_node
	}
	new_node.left = node
	node.parent = new_node
}

func (tree *RedBlackTree) rightRotate(node *redBlackNode) {
	new_node := node.left
	node.left = new_node.right
	if new_node.right != nil {
		new_node.right.parent = node
	}
	new_node.parent = node.parent
	if node.parent == nil {
		tree.root = new_node
	} else if node == node.parent.right {
		node.parent.right = new_node
	} else {
		node.parent.left = new_node
	}
	new_node.right = node
	node.parent = new_node
}

func (tree *RedBlackTree) Search(value int) bool {
	if searchNode(tree.root, value) == nil {
		return false
	} else {
		return true
	}
}

func searchNode(node *redBlackNode, value int) *redBlackNode {
	if node == nil || node.value == value {
		return node
	}
	if value < node.value {
		return searchNode(node.left, value)
	}
	return searchNode(node.right, value)
}

func RedBlackTreeFromArray(values []int) *RedBlackTree {
	tree := NewRedBlackTree()
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

// Scapegoat Tree
type scapegoatNode struct {
	Value int
	Left  *scapegoatNode
	Right *scapegoatNode
	Size  int
}

type ScapegoatTree struct {
	Alpha            float64
	Root             *scapegoatNode
	Size             int
	MaxSize          int
	TreeIsUnbalanced chan bool
}

func NewScapegoatTree(alpha float64) *ScapegoatTree {
	if alpha < 0.5 || alpha > 1.0 {
		panic(errors.New("Alpha is out of range. It should be between 0.5 and 1.0."))
	}
	return &ScapegoatTree{
		Alpha:            alpha,
		Root:             nil,
		Size:             0,
		MaxSize:          0,
		TreeIsUnbalanced: make(chan bool),
	}
}

func (tree *ScapegoatTree) Insert(value int) {
	newNode := &scapegoatNode{
		Value: value,
		Left:  nil,
		Right: nil,
		Size:  1,
	}

	if tree.Root == nil {
		tree.Root = newNode
	} else {
		tree.insertNode(tree.Root, newNode)
	}

	tree.Size++
	tree.MaxSize = int(math.Max(float64(tree.Size), float64(tree.MaxSize)))

	if tree.needsRebuild() {
		go func() {
			tree.TreeIsUnbalanced <- true
		}()
	}
}

func (tree *ScapegoatTree) insertNode(parent *scapegoatNode, newNode *scapegoatNode) {
	if newNode.Value < parent.Value {
		if parent.Left == nil {
			parent.Left = newNode
		} else {
			tree.insertNode(parent.Left, newNode)
		}
	} else {
		if parent.Right == nil {
			parent.Right = newNode
		} else {
			tree.insertNode(parent.Right, newNode)
		}
	}

	parent.Size = 1 + tree.size(parent.Left) + tree.size(parent.Right)
}

func (tree *ScapegoatTree) Search(value int) bool {
	if tree.searchNode(tree.Root, value) == nil {
		return false
	} else {
		return true
	}
}

func (tree *ScapegoatTree) searchNode(node *scapegoatNode, value int) *scapegoatNode {
	if node == nil || value == node.Value {
		return node
	}

	if value < node.Value {
		return tree.searchNode(node.Left, value)
	} else {
		return tree.searchNode(node.Right, value)
	}
}

func (tree *ScapegoatTree) size(node *scapegoatNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func (tree *ScapegoatTree) needsRebuild() bool {
	if tree.Root == nil {
		return false
	}
	return float64(tree.Size) > tree.Alpha*float64(tree.MaxSize)
}

func (tree *ScapegoatTree) Delete(value int) {
	tree.Root = tree.deleteNode(tree.Root, value)
}

func (tree *ScapegoatTree) deleteNode(node *scapegoatNode, value int) *scapegoatNode {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = tree.deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = tree.deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		node.Value = tree.minValue(node.Right)
		node.Right = tree.deleteNode(node.Right, node.Value)
	}

	node.Size = 1 + tree.size(node.Left) + tree.size(node.Right)
	return node
}

func (tree *ScapegoatTree) minValue(node *scapegoatNode) int {
	min := node
	for min.Left != nil {
		min = min.Left
	}
	return min.Value
}

func ScapegoatTreeFromArray(alpha float64, values []int) *ScapegoatTree {
	tree := NewScapegoatTree(alpha)
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func main() {
	array_generators := []func(int) []int{generateSortedArray, generateReversedSortedArray, generateRandomSortedArray, generateAlmostSortedArray}
	// sizes := []int{10, 100, 1000, 10000, 100000, 1000000}
	sizes := []int{10, 100, 1000}

	for _, array_generator := range array_generators {
		for _, size := range sizes {
			fmt.Println("Array size:", size, "Array generator:", array_generator)
			array := array_generator(size)

			bst := BinarySearchTreeFromArray(array)
			bst.Search(10)

			avl := AVLTreeFromArray(array)
			avl.Search(10)

			rbt := RedBlackTreeFromArray(array)
			rbt.Search(10)

			sgt := ScapegoatTreeFromArray(0.6, array)
			sgt.Search(10)
		}
	}
}
