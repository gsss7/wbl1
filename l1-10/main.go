package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.5, -0.5}
	groups := make(map[int][]float32)

	for _, temp := range temperatures {
		group := int(temp/10) * 10
		groups[group] = append(groups[group], temp)
	}
	fmt.Println(groups)
}
