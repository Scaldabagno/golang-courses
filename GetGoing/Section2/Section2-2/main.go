package main

import (
	"fmt"
	"strings"
)

func main() {

	// var m1 string
	// m1 = "my name"

	m1 := "my name"
	// m1 := "shsh" //error
	// m1 = "shsh"  //ok
	// m2 := "name"
	// fmt.Println(strings.Contains(m1, m2)) //true
	// fmt.Println(strings.ReplaceAll(m1, "m", "NO"))
	fmt.Println(strings.Split(m1, " "))
}
