package main

import "fmt"

// Anything prints anything
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	Anything(2.44)       // 2.44
	Anything("angad")    // angad
	Anything(struct{}{}) // {}

	mymap := make(map[string]interface{})

	mymap["name"] = "nIUHSHBCI"
	mymap["age"] = 10
	fmt.Println(mymap)
	// map[age:10 name:nIUHSHBCI]
}
