package main

import "fmt"

func main() {
	//  simple iteration over a range
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over a collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// using for as while
	i := 0
	for i < 5 {
		fmt.Println("While", i)
		i++
	}
	// break and continue statements

}
