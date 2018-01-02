package main

import "fmt"

func main() {
	slice1 := []int{2, 3, 4, 5, 6, 7, 8,}
	oddNums := filter(slice1, isOdd)
	evenNums := filter(slice1, isEven)
	fmt.Printf("%v\n", oddNums)
	fmt.Printf("%v\n", evenNums)
	testSqrt(4)
}
