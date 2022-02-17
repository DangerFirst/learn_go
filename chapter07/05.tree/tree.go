package main

import "fmt"

type Node struct {
	date  int
	root  *Node
	left  *Node
	right *Node
}

func buildTree() *Node {
	n1 := &Node{date: 51}
	n2 := &Node{date: 35}
	n3 := &Node{date: 65}

	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1

	return n1
}

func insertNode(root, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if newNode.date == root.date {
		return root
	}
	if newNode.date < root.date {
		if root.left == nil {
			//找到位置，插入数据
			root.left = newNode
			newNode.root = root
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			//找到位置，插入数据
			root.right = newNode
			newNode.root = root
		} else {
			insertNode(root.right, newNode)
		}
	}
	return root
}

func deleteLeaf(root *Node, data int) *Node {
	leftRoot := root
	if leftRoot.date == data && leftRoot.left == nil && leftRoot.right == nil {
		leftRoot = leftRoot.root
		right := root
		if leftRoot.left == right {
			//删除左边的叶子
			leftRoot.left = nil
			right.root = nil
			return leftRoot
		} else {
			//删除右边的叶子
			leftRoot.right = nil
			right.root = nil
			return leftRoot
		}
	}
	return root
}

func deleteNode(root *Node, data int) *Node {
	if data < root.date {
		deleteNode(root.left, data)
	} else if data > root.date {
		deleteNode(root.right, data)
	} else {
		//现在root指向要删除的节点
		leftNextGen := findNextGenFromLeft(root.left)
		rightNextGen := findNextGenFromRight(root.right)
		if leftNextGen == nil && rightNextGen == nil {
			//现在要删除的是叶子节点，即最底部节点
			top := root.root
			down := root
			if top.left == down {
				//表示是左子树
				top.left = nil
				down.root = nil
				return nil
			} else {
				//表示是右子树
				top.right = nil
				down.root = nil
				return nil
			}
		} else if leftNextGen != nil {
			root.date = leftNextGen.date
			deleteNode(leftNextGen, leftNextGen.date)
		} else if rightNextGen != nil {
			root.date = rightNextGen.date
			deleteNode(rightNextGen, rightNextGen.date)
		}
	}
	return root
}

func findNextGenFromLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

func findNextGenFromRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}

func main() {
	root := buildTree()
	insertNode(root, &Node{date: 43})
	insertNode(root, &Node{date: 28})
	fmt.Println("done")

	fmt.Println("删除Node")
	deleteNode(root, 43)
	fmt.Println("删除结束")
}
