package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{[]int{}, 5, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 3, -1},
	}

	for _, test := range tests {
		result := binarySearch(test.arr, test.target)
		status := "✓"
		if result != test.expected {
			status = "✗"
		}
		fmt.Printf("%s Search %d in %v: got %d, expected %d\n",
			status, test.target, test.arr, result, test.expected)
	}
}
