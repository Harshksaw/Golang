package main

import "fmt"


//slices - // Slices are a data structure in Go that allows you to work with a sequence of elements.
// They are similar to arrays, but they are more flexible and can grow and shrink in size.
// Slices are built on top of arrays, and they provide a more convenient way to work with sequences of elements.
// Slices are a reference type, which means that when you pass a slice to a function, you are passing a reference to the underlying array.

func main() {

	nums := make([]int, 2)

	//var nums []int
	fmt.Println(nums)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)


	fmt.Println(len(nums))

	for _, num:= range nums {
		fmt.Println(num)
	}
	for i,e:= range "Harsh" {
		fmt.Println(i, string(e))

	}

}
