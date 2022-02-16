package main

import "fmt"

type LinkNode struct {
	date int
	next *LinkNode
}

func main() {
	n1 := &LinkNode{
		date: 1,
		next: nil,
	}
	n2 := &LinkNode{
		date: 2,
		next: nil,
	}
	n3 := &LinkNode{
		date: 3,
		next: nil,
	}
	n4 := &LinkNode{
		date: 4,
		next: nil,
	}
	n6 := &LinkNode{
		date: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	{
		rangeLink(n1)
	}

	{
		fmt.Println("插入节点")
		n5 := &LinkNode{
			date: 5,
			next: nil,
		}
		n1 = insertNode(n1, n5)
		n1 = insertNode(n1, &LinkNode{
			date: 7,
			next: nil,
		})
		n1 = insertNode(n1, &LinkNode{
			date: 3,
			next: nil,
		})
		n1 = insertNode(n1, &LinkNode{
			date: 0,
			next: nil,
		})
		rangeLink(n1)
	}
	{
		fmt.Println("删除节点")
		n1 = deleteNode(n1, 0)
		n1 = deleteNode(n1, 3)
		n1 = deleteNode(n1, 7)
		n1 = deleteNode(n1, 8)
		rangeLink(n1)
	}
}

func deleteNode(root *LinkNode, date int) *LinkNode {
	tempNode := root
	if root != nil && root.date == date {
		if root.next == nil {
			return nil
		}
		rootNext := root.next
		root.next = nil
		return rootNext
	}
	for {
		if tempNode.next == nil {
			break
		}
		nextNode := tempNode.next
		if nextNode.date == date {
			tempNode.next = nextNode.next
			nextNode.next = nil
			return root
		}
		tempNode = nextNode
	}
	return root
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tempNode := root
	if root.date > newNode.date {
		newNode.next = root
		return newNode
	}
	for {
		if tempNode != nil {
			if newNode.date > tempNode.date {
				if tempNode.next == nil {
					tempNode.next = newNode
				} else {
					if tempNode.next.date >= newNode.date {
						newNode.next = tempNode.next
						tempNode.next = newNode
						break
					}
				}
			}
		} else {
			break
		}
		tempNode = tempNode.next
	}
	return root
}

func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.date)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}
