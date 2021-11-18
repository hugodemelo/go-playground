package main

import "fmt"

func myMap[T any](arr *[]T, fn func(T) T) []T {
	result := []T{}

	for _, value := range *arr {
		result = append(result, fn(value))
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3}

	fmt.Println(myMap(&numbers, func(i int) int {
		return i * 2
	}))
}
