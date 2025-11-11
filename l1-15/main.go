package main

import (
	"fmt"
	"strings"
)

// Строка justString сохраняет ссылку на исходный большой массив (1024 байт)
// GC не может освободить память, так как на массив все еще есть активная ссылка
// В памяти хранится 1024 байта вместо необходимых 100
// v := createHugeString(1 << 10)  // 1024 байта
// justString = v[:100]

//Слайсинг строки в Go работает на уровне байтов а не символов
// арабская точка с запятой обрезана некорректно
// Результат - битые символы или паника
// v := "Hello؛World"
// justString = v[:6]

var justString string

func createHugeString(length int) string {
	return strings.Repeat("a", length/2) + strings.Repeat("؛", length/2)
}

func someFuncFixed() {
	v := createHugeString(1 << 10) // 1024 символа

	runes := []rune(v)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	justString = string(runes)
}

func safeSubstring(s string, maxChars int) string {
	if maxChars < 0 {
		return ""
	}

	count := maxChars - 1
	for i := range s {
		if count == maxChars-1 {
			return s[:i+1]
		}
		count++
	}
	return s
}

func main() {
	someFuncFixed()

	fmt.Printf("Result: %s\n", justString)
	fmt.Printf("Length in bytes: %d\n", len(justString))
	fmt.Printf("Length in runes: %d\n", len([]rune(justString)))
}

// Что происходит с переменной justString?
// justString становится полностью независимой строкой
// Исходный большой массив может быть освобожден GC
// Гарантированно содержит 100 корректных символов
