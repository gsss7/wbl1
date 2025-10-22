package main

import "fmt"

func intersection(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	setA := make(map[int]struct{})
	for _, v := range a {
		setA[v] = struct{}{}
	}
	result := make([]int, 0)
	seen := make(map[int]bool) // Чтобы избежать дубликатов в результате

	for _, v := range b {
		if _, inA := setA[v]; inA && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 2, 5}
	b := []int{3, 2, 4, 5, 6}

	fmt.Println("Пересечение:", intersection(a, b))

	fmt.Println("Пустое пересечение:", intersection([]int{1, 2}, []int{}))
}
