package main

import "fmt"

func main() {

	// var m1 int
	// m1 = 2
	// fmt.Println(m1)

	// var (
	// 	m1 = 2
	// 	m2 = 3
	// )
	// fmt.Println(m1 + m2)

	var m1 int32
	var m2 int64
	//fmt.Println(m1 + m2) //error
	fmt.Println(int64(m1) + m2)
}

// atomic datatypes
// string
// int
// int32
// int64
// uint
// uint32
// uint64

//unsafe
// Pointers

//abstract data types
//map[]<datatype>
// struct{}
// interface{}
