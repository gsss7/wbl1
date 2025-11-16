package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var sort func(int, int)
	sort = func(low, high int) {
		if low >= high {
			return
		}

		pivotIndex := partition(arr, low, high)

		sort(low, pivotIndex-1)
		sort(pivotIndex+1, high)
	}

	sort(0, len(arr)-1)
}

// Функция разделения массива
func partition(arr []int, low, high int) int {
	pivotIndex := (low + high) / 2 // средний элемент

	pivot := arr[pivotIndex]

	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i
}

func quickSortImmutable(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var left, right []int

	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	result := append(quickSortImmutable(left), pivot)
	result = append(result, quickSortImmutable(right)...)

	return result
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{5, 4, 3, 2, 1, 10, 22, 9, 6, 6, 50, 0, 3, 5}
	fmt.Println("Original:", arr1)
	quickSort(arr1)
	fmt.Println("Sorted (in-place):", arr1)
	fmt.Println("Is sorted:", isSorted(arr1))

	fmt.Println()

	arr2 := []int{5, 4, 3, 2, 1, 10, 22, 9, 6, 6, 50, 0, 3, 5}
	fmt.Println("Original:", arr2)
	sorted := quickSortImmutable(arr2)
	fmt.Println("Sorted (immutable):", sorted)
	fmt.Println("Is sorted:", isSorted(sorted))

	fmt.Println()

	largeArr := make([]int, 1000)
	for i := range largeArr {
		largeArr[i] = rand.Intn(10000)
	}

	quickSort(largeArr)
	fmt.Println("Large array sorted:", isSorted(largeArr))
}
