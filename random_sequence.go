package main

import (
	"fmt"
	"math/rand"
)

func random_sequence(n int) []int {
	sequential := make([]int, n, n)
	for i := 0; i < n; i++ {
		sequential[i] = i
	}
	rands := n_random_ints(n)
	_, random_sequence := MergeSortDuplicated(rands, sequential)
	return random_sequence
}

func MergeSortDuplicated(slice, slice_sequential []int) ([]int, []int) {
	//sorts the first slice performing the same operations on the second

	if len(slice) < 2 {
		return slice, slice_sequential
	}
	mid := len(slice) / 2

	left, left_sequential := MergeSortDuplicated(slice[:mid], slice_sequential[:mid])
	right, right_sequential := MergeSortDuplicated(slice[mid:], slice_sequential[mid:])
	return MergeDuplicated(left, right, left_sequential, right_sequential)
}

func MergeDuplicated(left, right, left_sequential, right_sequential []int) ([]int, []int) {
	//merges the first 2 slices to be ordered, duplicating the operation to the second two
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	slice_sequential := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			slice_sequential[k] = right_sequential[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			slice_sequential[k] = left_sequential[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			slice_sequential[k] = left_sequential[i]
			i++
		} else {
			slice[k] = right[j]
			slice_sequential[k] = right_sequential[j]
			j++
		}
	}
	return slice, slice_sequential
}

func n_random_ints(n int) []int { //probably a lib function for this but i did with no internet
	rands := make([]int, n, n)
	for i := 0; i < n; i++ {
		rands[i] = rand.Intn(10000)
	}
	return rands
}

func main() {
	n := 8
	fmt.Println(random_sequence(n))
}
