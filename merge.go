// Write a function that merges
// two sorted lists into a new sorted list,
// e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].

// https://github.com/smcguire56/GoLangFundamentals :: 2017-10-05
// Sean McGuire
package main

import "fmt"

func main() {
	list1 := []byte{7, 1, 8, 47, 4, 5}
	list2 := []byte{8, 55, 11, 47, 6, 13}

	fmt.Println(merge(list1, list2))
}

func merge(list1 []byte, list2 []byte) []byte {
    return append(list1, list2...)
}