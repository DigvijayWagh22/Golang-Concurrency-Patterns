package main

import "fmt"

func byValue(val int) {
	val = 0
}

func byPtr(ptr *int) {
	*ptr = 0
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, elem := range nums {
		sum += elem
	}
	fmt.Println("sum is ", sum)

	for idx, num := range nums {
		if num == 3 {
			fmt.Println("Index of 3 is ", idx)
		}
	}

	store := map[string]int{"a": 1, "b": 0, "e": 8}
	for key, value := range store {
		fmt.Println(key, " -> ", value)
	}

	for pair := range store {
		fmt.Println(pair)
	}

	for i, c := range "Hello world" {
		fmt.Println("the idx is ", i, " the value is ", c)
	}

	// pointer section

	i := 1
	fmt.Println("Initial value ", i)

	byValue(i)
	fmt.Println("Passing by value ", i)
	byPtr(&i)
	fmt.Println("Passing by ptr ", i)

	fmt.Println("pointer value ", &i)

}
