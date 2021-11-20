package generics

import (
	"fmt"
	"strconv"
)

func myMap[T any, U any](arr *[]T, fn func(T) U) []U {
	result := []U{}

	for _, value := range *arr {
		result = append(result, fn(value))
	}
	return result
}

func RunGenerics() {
	numbers := &[]int{1, 2, 3}

	fmt.Println(myMap(numbers, func(i int) int {
		return i * 2
	}))
	
	fmt.Println(myMap(numbers, func(i int) string {
		return strconv.Itoa(i)
	}))
}
