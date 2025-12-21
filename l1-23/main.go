package main

import (
	"fmt"
	"runtime"
)

func removeElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	copy(slice[index:], slice[index+1:])
	var zero T
	slice[len(slice)-1] = zero
	return slice[:len(slice)-1]
}

func removeElementSafe[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	result := make([]T, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}

type BigObject struct {
	data [1 << 20]byte
}

func printMemUsage(description string) {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Занято %v MB\n", description, m.Alloc/1024/1024)
}

func main() {
	fmt.Println("=== Демонстрация удаления с предотвращением утечек ===")
	objects := make([]*BigObject, 5)
	for i := range objects {
		objects[i] = &BigObject{}
		objects[i].data[0] = byte(i + 1)
	}
	printMemUsage("После создания 5 объектов")
	removedObject := objects[2]
	objects = removeElement(objects, 2)

	fmt.Printf("\nУдалён объект с данными: %v\n", removedObject.data[0])
	fmt.Printf("Длина слайса после удаления: %d\n", len(objects))

	runtime.GC()
	printMemUsage("После удаления и GC")

	fmt.Println("\n=== Пример с простыми значениями ===")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Исходный: %v\n", numbers)
	numbers = removeElement(numbers, 4)
	fmt.Printf("После removeElement(4): %v\n", numbers)

	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers2 = removeElementSafe(numbers2, 4)
	fmt.Printf("После removeElementSafe(4): %v\n", numbers2)
	fmt.Println("\n=== Граничные случаи ===")
	testSlice := []string{"a", "b", "c"}
	testSlice = removeElement(testSlice, 10)
	fmt.Printf("Некорректный индекс: %v\n", testSlice)
	testSlice = removeElement(testSlice, 0)
	fmt.Printf("После удаления первого: %v\n", testSlice)
	testSlice = removeElement(testSlice, len(testSlice)-1)
	fmt.Printf("После удаления последнего: %v\n", testSlice)
}
