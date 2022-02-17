package main

import "fmt"

type LinkNode struct {
	date     int
	next     *LinkNode
	previous *LinkNode
}

func buildLink() *LinkNode {
	n1 := &LinkNode{
		date: 1,
	}
	n2 := &LinkNode{
		date: 2,
	}
	n3 := &LinkNode{
		date: 3,
	}
	n1.next = n2
	n2.previous = n1
	n2.next = n3
	n3.previous = n2
	return n1
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tempNode := root
	//列表是空的情况新增
	if root == nil {
		return newNode
	}
	//在列表头新增
	if root.date >= newNode.date {
		newNode.next = root
		root.previous = newNode
		return newNode
	}
	for {
		if tempNode.next == nil {
			//	已经到头
			tempNode.next = newNode
			newNode.previous = tempNode
			return root
		} else {
			if tempNode.next.date >= newNode.date {
				newNode.previous = tempNode
				newNode.next = tempNode.next
				tempNode.next = newNode
				newNode.next.previous = newNode
				return root
			}
		}
		tempNode = tempNode.next
	}
}

func deleteNode(root *LinkNode, date int) *LinkNode {
	if root == nil || root.next == nil && root.date == date {
		return nil
	}
	if root.next != nil && root.date == date {
		//要删除的数据在第一个节点
		rootNext := root.next
		root.next.previous = nil
		root.next = nil
		return rootNext
	}
	tmpNode := root
	for {
		if tmpNode.next == nil {
			//找到链表的尾部，仍然没有找到要删除的数据，直接返回原root
			return root
		} else {
			if tmpNode.next.date == date {
				//找到节点，开始删除，删除后返回原节点
				nextNode := tmpNode.next
				tmpNode.next = nextNode.next
				//删尾的时候需要判断下
				if nextNode.next != nil {
					nextNode.next.previous = tmpNode
				}
				nextNode.previous = nil
				nextNode.next = nil
				return root
			}
		}
		tmpNode = tmpNode.next
	}
}

func RangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tempNode := root
	for {
		fmt.Println(tempNode.date)
		if tempNode.next == nil {
			break
		}
		tempNode = tempNode.next
	}
	fmt.Println("从尾到头")
	for {
		fmt.Println(tempNode.date)
		if tempNode.previous == nil {
			break
		}
		tempNode = tempNode.previous
	}
}

func main() {
	root := buildLink()
	RangeLink(root)
	root = insertNode(root, &LinkNode{date: 3})
	root = insertNode(root, &LinkNode{date: 7})
	root = insertNode(root, &LinkNode{date: 0})
	fmt.Println("新增后")
	RangeLink(root)
	root = deleteNode(root, 3)
	root = deleteNode(root, 0)
	root = deleteNode(root, 7)
	fmt.Println("删除后")
	RangeLink(root)

}
