package main

import "fmt"

func createSet(world []string) []string {
	set := make(map[string]struct{})

	for _, word := range world {
		set[word] = struct{}{}
	}

	result := make([]string, 0, len(set))
	for word := range set {
		result = append(result, word)
	}
	return result
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree", "dog"}
	uniqueWords := createSet(words)

	fmt.Println("Множество уникальных слов:", uniqueWords)
}
