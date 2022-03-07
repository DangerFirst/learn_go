package main

import (
	"fmt"
	"time"
)

type Node struct {
	data  int
	root  *Node
	left  *Node
	right *Node
}

func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}

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
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
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
	if leftRoot.data == data && leftRoot.left == nil && leftRoot.right == nil {
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
	if data < root.data {
		deleteNode(root.left, data)
	} else if data > root.data {
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
			root.data = leftNextGen.data
			deleteNode(leftNextGen, leftNextGen.data)
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data)
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

var totalCompare int = 0

func main() {
	//mainForPreSearch()
	//mainForMidSearch()
	mainForPostSearch()
}

func mainForPreSearch() {
	var root *Node
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束", buildFinishTime.Sub(startTime))
	for i := 0; i < 10000; i++ {
		preSearch(root, 11000)
		preSearch(root, 3016)
		preSearch(root, 2)
		preSearch(root, 44)
	}
	finishTime := time.Now()
	fmt.Println("总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))

}

func mainForMidSearch() {
	var root *Node
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束", buildFinishTime.Sub(startTime))
	for i := 0; i < 10000; i++ {
		midSearch(root, 11000)
		midSearch(root, 3016)
		midSearch(root, 2)
		midSearch(root, 44)
	}
	finishTime := time.Now()
	fmt.Println("总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))

}

func mainForPostSearch() {
	var root *Node
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertNode(root, &Node{data: int(v)})
	}
	buildFinishTime := time.Now()
	fmt.Println("构建结束", buildFinishTime.Sub(startTime))
	for i := 0; i < 10000; i++ {
		postSearch(root, 11000)
		postSearch(root, 3016)
		postSearch(root, 2)
		postSearch(root, 44)
	}
	finishTime := time.Now()
	fmt.Println("总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))

}
