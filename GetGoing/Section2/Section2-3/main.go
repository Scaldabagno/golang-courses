package main

import "fmt"

func todo() {
	// varr arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}

	fmt.Println(arr)
	arr2 = append(arr2, "is", "angad")
	fmt.Println(arr2)
}

func main() {
	todo()
}
