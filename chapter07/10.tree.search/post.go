package main

func postSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	if root.data > targetNum {
		if postSearch(root.left, targetNum) {
			return true
		}
	}
	if root.data < targetNum {
		if postSearch(root.right, targetNum) {
			return true
		}
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	return false
}
