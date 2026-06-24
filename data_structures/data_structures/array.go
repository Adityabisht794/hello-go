package main

import "fmt"

func main() {
	// Array declaration
	var nums [5]int
	fmt.Println(nums)

	// Array manipulation
	nums[0] = 6
	nums[1] = 7
	fmt.Println(nums)

	// length of array
	fmt.Println(len(nums))

	// Capacity of array
	fmt.Println(cap(nums))

	// Dynamic array declaration
	var da []int
	da = append(da, 1, 1, 3, 4, 5, 6, 7, 7, 8, 9, 10)
	fmt.Println(da)

	fmt.Println(cap(da))
	fmt.Println(len(da))
	da = append(da, 14, 15)
	fmt.Println(cap(da))
	fmt.Println(len(da))

	// Slice ie similar to vector and make function.

	var v= make([]int , 3 )
	fmt.Println(v)

	// 2d slice declaration (slice of arrays) so we can append rows
	var matrix [][]int
	matrix = append(matrix, []int{1, 2, 3, 3}, []int{4, 5, 6, 7} )

	fmt.Println(matrix) 
}
