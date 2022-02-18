package main

func mergeSort(arr *[]int) {

}

func mergeArr(left, right []int) []int {
	out := []int{}
	leftI, rightI := 0, 0
	for {
		if leftI == len(left) || rightI == len(right) {
			break
		}
		if left[leftI] < right[rightI] {
			out = append(out, left[leftI])
			leftI++
			continue
		} else {
			out = append(out, right[rightI])
			rightI++
			continue
		}
	}
	for ; left[leftI] < len(left); leftI++ {
		out = append(out, left[leftI])
	}
	for ; left[rightI] < len(right); rightI++ {
		out = append(out, right[rightI])
	}
	return out
}
