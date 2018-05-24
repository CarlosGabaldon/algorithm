package main

import (
	"fmt"
	"sort"
	"strings"
)

func SecondHighest(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-2]

}
func main() {
	numbers := []int{2, 4, 1, 60, 3, 40}
	fmt.Println("Given the numbers: ", numbers)
	fmt.Println("Find the second highest number")
	secondHighest := SecondHighest(numbers)
	fmt.Println("Second highest number: ", secondHighest)
	fmt.Println(strings.ToUpper("that was fun!"))
}
