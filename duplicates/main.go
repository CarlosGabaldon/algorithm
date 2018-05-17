package main

import "fmt"

/*
hasdups is a linear solution to look  for duplicate numbers
within an array of integers. It is O(N) for time complexity and O(N)
for space complexity. It has a limitation of only handling numbers less 1000,
this is becuase I did not know how to create an array in Go with dynamic
capacity.
*/
func hasDups(array []int) bool {
	steps := 0
	existingNumberMap := make([]int, 10, 1000)
	var number int

	for i := 0; i < len(array); i++ {
		steps++

		number = array[i]

		if number > len(existingNumberMap) {
			resize := number + 1
			existingNumberMap = existingNumberMap[:resize]
		}

		if existingNumberMap[number] == 0 {
			existingNumberMap[number] = 1
		} else {

			printSteps(steps)

			return true
		}

	}

	printSteps(steps)
	return false

}

func printSteps(steps int) {
	fmt.Printf("Using a Linear search with time complexity O(N) took %d steps\n", steps)

}

func main() {
	array := []int{2, 3, 7, 4, 11, 100, 99, 12, 34, 899}
	fmt.Println("Given the array:")
	fmt.Println(array)
	fmt.Println("Find any duplicate numbers")
	dups := hasDups(array)
	fmt.Println("Array has duplication of values: ", dups)
}
